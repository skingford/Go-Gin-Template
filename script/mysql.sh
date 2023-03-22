#!/bin/bash
container_name=mysql5.7
file_path=$HOME/podman/mysql/mysql5.7

function start()
{
	podman run -itd --rm \
		-p 3306:3306 \
    -e MYSQL_USER=admin \
    -e MYSQL_PASSWORD=123456 \
		-e MYSQL_ROOT_PASSWORD=root \
		-v $file_path/data:/var/lib/mysql:rw \
		-v $file_path/log:/var/log/mysql:rw \
		-v $file_path/conf/my.cnf:/etc/mysql/my.cnf:rw \
		-v /etc/localtime:/etc/localtime:ro \
		--name $container_name \
		mysql:5.7
}

function debug()
{
	podman run -it --rm \
		-p 3306:3306 \
		-e MYSQL_ROOT_PASSWORD=root \
		-v $file_path/data:/var/lib/mysql:rw \
		-v $file_path/log:/var/log/mysql:rw \
		-v $file_path/conf/my.cnf:/etc/mysql/my.cnf:rw \
		-v /etc/localtime:/etc/localtime:ro \
		--name $container_name \
		mysql
}

function stop()
{
	podman stop $container_name
}

function podman_ps()
{
	podman ps -a
}

function login()
{
	podman exec -it $container_name bash
}

function usage()
{
	echo -e "Usage: $0 start|debug|stop|ps|login"
	exit 0
}


if [[ $1 == "start" ]];then
	start
elif [[ $1 == "debug" ]];then
	debug
elif [[ $1 == "stop" ]];then
	stop
elif [[ $1 == "ps" ]];then
	podman_ps
elif [[ $1 == "login" ]];then
	login
else
	usage
fi