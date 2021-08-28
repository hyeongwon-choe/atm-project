# Requirement
* Linux environment
* gcc
# Install
``` bash
git clone http://github.com/hyeongwon-choe/atm-project
./installGoBinary.sh
```
# Test
``` bash
./test.sh
```
* You can find testcode at atm/atm_test.go
# How to integrate with a real bank system
* You must implement dao that can interact real bank data
  * Implement interface of dao/dao.go
  * You can refer to dao/dummy_dao/dummy_dao.go
* Inject your dao to Atm object, and use it
  * You can refer to atm/atm_test.go:Test_ATM_Basic_Scenario
