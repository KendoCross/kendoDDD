package testcmd

import "helm.sh/helm/v3/pkg/release"

type helmEntity struct {
}

func newHelmEnByOV() *helmEntity {
	return new(helmEntity)
}

func (en *helmEntity) ReleaseList() (rls []*release.Release, err error) {
	rls, err = helmClient.FindReleases()
	return
}

func (en *helmEntity) ReleaseUpd() (err error) {
	helmClient.HelmUpgrade("", "", "", "", nil)
	return
}
