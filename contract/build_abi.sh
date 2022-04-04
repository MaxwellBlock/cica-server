project=cica-contract
network=bsctest
env=develop
version=latest
key=abi
contracts=(Hello)

for (( i=0;i<${#contracts[@]};i++))
do
  # shellcheck disable=SC2006
  # shellcheck disable=SC2019
  # shellcheck disable=SC2018
  small=`echo "${contracts[i]}" | tr 'A-Z' 'a-z'`

  curl "http://47.241.117.7:2022/contracts/get?project=${project}&env=${env}&version=${version}&key=${key}&network=${network}&contract=${contracts[i]}" > "${contracts[i]}".json
  if [ ! -d "${small}" ];then
    mkdir "${small}"
  fi

  abigen --abi="${contracts[i]}".json --pkg="${small}" --out=./"${small}"/"${contracts[i]}".go
done;
