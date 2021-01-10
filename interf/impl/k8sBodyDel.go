package impl

import "menu/db"

func (k K8sBodyList) K8sBodyDel(address string) {
	db.Del([]byte(address), db.K8sList)
}
