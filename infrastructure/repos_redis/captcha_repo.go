package repos_redis

import "time"

//活动账号基本信息DBKey
const ConstKeyPre = "CaptchaRepoRedisDBKey"

type captchaRepo struct{}

func NewcaptchaRepo() *captchaRepo {
	return new(captchaRepo)
}

func (repo *captchaRepo) SetCaptcha(phone, code string) (err error) {
	_, err = driver.Set(ConstKeyPre+phone, code, time.Minute*10)
	return
}
func (repo *captchaRepo) GetCaptcha(phone string) (code string, err error) {
	code, err = driver.Get(ConstKeyPre + phone)
	return
}

func (repo *captchaRepo) DelCaptcha(phone string) (code int64, err error) {
	code, err = driver.Del(ConstKeyPre + phone)
	return
}
