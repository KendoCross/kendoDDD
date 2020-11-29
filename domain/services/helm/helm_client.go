package helm

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ghodss/yaml"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/release"
)

const (
	HelmKubeConfTmpFilePath = "/tmp/"
)

const (
	DefaultReadTimeout  = 1
	DefaultWriteTimeout = 2
)

type ReleaseInfo struct {
	Name              string    `json:"name"`
	Version           string    `json:"version"`
	Values            string    `json:"values"`
	Status            string    `json:"status"`
	DeployedTime      time.Time `json:"deployedTime"`      // current release version deployed time
	FirstDeployedTime time.Time `json:"firstDeployedTime"` // the release first deployed time
}

type Options struct {
	// timeout = WriteTimeout * time.Minute
	WriteTimeout time.Duration
	// timeout = ReadTimeout * time.Minute
	ReadTimeout time.Duration
}

type HelmAction struct {
	Options
	namespace    string
	kubeConfig   string
	actionConfig *action.Configuration
}

func NewHelmAction(kubeconfFile, ns string, options ...*Options) (*HelmAction, error) {
	ac, err := getHelmActionConf(kubeconfFile, ns)
	if err != nil {
		return nil, err
	}
	var readTimeout, writeTimeout time.Duration
	if len(options) >= 1 {
		readTimeout = options[0].ReadTimeout
		writeTimeout = options[0].WriteTimeout
	} else {
		readTimeout = DefaultReadTimeout
		writeTimeout = DefaultWriteTimeout
	}

	return &HelmAction{
		Options: Options{
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
		namespace:    ns,
		kubeConfig:   kubeconfFile,
		actionConfig: ac,
	}, nil
}

func (a *HelmAction) HelmInstall(name, namespace, version, chartPath string, vals map[string]interface{}) error {
	chart, err := loader.LoadDir(chartPath)
	if err != nil {
		return err
	}
	chart.Metadata.Version = version

	hClient := action.NewInstall(a.actionConfig)
	hClient.Namespace = namespace
	hClient.ReleaseName = name
	hClient.Version = version
	hClient.Timeout = time.Minute * a.WriteTimeout

	if err := isChartInstallable(chart); err != nil {
		return err
	}
	// 发送请求
	_, err = hClient.Run(chart, vals)
	if err != nil {
		return err
	}
	return nil
}

func (a *HelmAction) HelmUpgrade(name, namespace, version, chartPath string, vals map[string]interface{}) error {
	chart, err := loader.LoadDir(chartPath)
	if err != nil {
		return err
	}
	chart.Metadata.Version = version

	hClient := action.NewUpgrade(a.actionConfig)
	hClient.Namespace = namespace
	hClient.Version = version
	hClient.Timeout = a.WriteTimeout * time.Minute

	if err := isChartInstallable(chart); err != nil {
		return err
	}
	// 发送请求
	_, err = hClient.Run(name, chart, vals)
	if err != nil {
		return err
	}

	return nil
}

func (a *HelmAction) HelmRollBack(name, version string) error {
	// get history release
	historyRelease, err := a.getHistoryReleaseByVersion(name, version)
	if err != nil {
		return err
	}

	hClient := action.NewRollback(a.actionConfig)
	hClient.Version = historyRelease.Version
	hClient.Timeout = a.WriteTimeout * time.Minute

	// 发送请求
	if err := hClient.Run(name); err != nil {
		return err
	}

	return nil
}

func (a *HelmAction) HelmUninstall(name string) error {
	hClient := action.NewUninstall(a.actionConfig)
	hClient.Timeout = a.WriteTimeout * time.Minute

	// 发送请求
	if _, err := hClient.Run(name); err != nil {
		return err
	}

	return nil
}

func (a *HelmAction) GetHistoryReleaseInfo(name string, num ...int) ([]*ReleaseInfo, error) {
	mapToYaml := func(vals map[string]interface{}) (string, error) {
		// map to json
		jsonBytes, err := json.Marshal(vals)
		if err != nil {
			return "", err
		}
		// json to yaml
		yamlBytes, err := yaml.JSONToYAML(jsonBytes)
		if err != nil {
			return "", err
		}
		return string(yamlBytes), nil
	}

	res, err := a.getHistoryReleaseList(name)
	if err != nil {
		return nil, err
	}
	releaseInfo := make([]*ReleaseInfo, 0)
	for _, rs := range res {
		var values string
		valStr, err := mapToYaml(mergeMaps(rs.Chart.Values, rs.Config))
		if err != nil {
			fmt.Printf(" this values invaliad: %v", err)
			values = fmt.Sprintf(" Invalid values: %v", err)
		} else {
			values = valStr
		}
		result := &ReleaseInfo{
			Name:              rs.Name,
			Version:           rs.Chart.Metadata.Version,
			Values:            values,
			Status:            string(rs.Info.Status),
			DeployedTime:      rs.Info.LastDeployed.Time,
			FirstDeployedTime: rs.Info.FirstDeployed.Time,
		}
		releaseInfo = append(releaseInfo, result)
	}

	if len(num) >= 1 {
		if num[0] <= len(releaseInfo) {
			return releaseInfo[len(releaseInfo)-num[0]:], nil
		}
	}
	return releaseInfo, nil
}

func (a *HelmAction) GetRelease(name string) (*release.Release, error) {
	hClient := action.NewGet(a.actionConfig)
	res, err := hClient.Run(name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *HelmAction) FindReleases() ([]*release.Release, error) {
	hClient := action.NewList(a.actionConfig)
	res, err := hClient.Run()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *HelmAction) getHistoryReleaseByVersion(name, version string) (*release.Release, error) {
	res, err := a.getHistoryReleaseList(name)
	if err != nil {
		return nil, err
	}
	for index, rs := range res {
		if rs.Chart.Metadata.Version == version {
			return res[index], nil
		}
	}
	return nil, fmt.Errorf(" Invaliad version %s, not exist", version)
}

func (a *HelmAction) getHistoryReleaseList(name string) ([]*release.Release, error) {
	hClient := action.NewHistory(a.actionConfig)
	res, err := hClient.Run(name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func isChartInstallable(ch *chart.Chart) error {
	switch ch.Metadata.Type {
	case "", "application":
		return nil
	}
	return fmt.Errorf("%s charts are not installable", ch.Metadata.Type)
}

func getHelmActionConf(kubeconfFile, ns string) (*action.Configuration, error) {
	if ns == "" {
		ns = "default"
	}

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(kubeconfFile, "", ns), ns, os.Getenv("HELM_DRIVER"), debug); err != nil {
		return nil, fmt.Errorf(" Init helm action conf failed: %v", err)
	}

	//genericclioptions.NewConfigFlags()
	return actionConfig, nil
}

func debug(format string, v ...interface{}) {
	format = fmt.Sprintf("[log-helm] %s\n", format)
	_ = log.Output(2, fmt.Sprintf(format, v...))
}

func mergeMaps(a, b map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = mergeMaps(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}
