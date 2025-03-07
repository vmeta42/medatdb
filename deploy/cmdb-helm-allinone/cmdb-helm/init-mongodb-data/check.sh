#!/bin/bash
#设置变量，url为你需要检测的目标网站的网址（IP或域名）

urlprefix=$1
#collectionName=$2
#importFileName=${collectionName}.json
#url=${urlprefix}/${collectionName}
#echo $url
#echo $importFileName

    #定义函数check_http：

    #使用curl命令检查http服务器的状态

    #-m设置curl不管访问成功或失败，最大消耗的时间为5秒，5秒连接服务为相应则视为无法连接

    #-s设置静>默连接，不显示连接时的连接速度、时间消耗等信息

    #-o将curl下载的页面内容导出到/dev/null(默认会在屏幕显示页面内容)

    #-w设置curl命令需要显示的内容%{http_code}，指定curl返回服务器的状态码

check_http(){

   status_code=$(curl   -m 5 -s -o /dev/null -w %{http_code} -H 'HTTP_BLUEKING_SUPPLIER_ID:0' -H 'BK_USER:migrate' $url)

}
cwd=$(cd "$(dirname "$0")"; pwd)
pushd $cwd
for collectionName in `find . -name "*.json"|sed 's#.*/##'`:
  do
    for i in `seq 1 60` :
      do
      collectionName=`echo ${collectionName//:/}`
      collection=`echo ${collectionName}|cut -d "." -f 1`
      url=${urlprefix}/${collection}
#      if [[ $collectionName == *:* ]];then
#          collectionName=${collectionName%:*}
#      fi

      echo $collectionName,$url
      check_http
      echo $status_code
      case $status_code in
        201)
          break
          ;;
        200)
          mongoimport --db cmdb  --file ${collectionName} -c ${collection} --uri="mongodb://cc:cc@${2}/?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=cmdb&authMechanism=SCRAM-SHA-256"
          ;;
        400)
          sleep 5
          ;;
        000)
          sleep 5
          ;;
       esac
done
done