#! /bin/bash
echo "============================"
echo "Cleanning all mysql/cnf dirs/files"
echo "============================"

# Stop mysql server
sudo runuser -l -c \
    '/home/chow/Softs/mysql/bin/mysqld_multi stop 3306'
sudo runuser -l -c \
    '/home/chow/Softs/mysql/bin/mysqld_multi stop 3307'
sudo runuser -l -c \
    '/home/chow/Softs/mysql/bin/mysqld_multi report'

sudo kill -9  `ps -ef|grep mysql|grep -v grep|awk '{print $2}'`

# delete sql dir
if [ -d "/home/chow/Softs/mysql" ];then
    echo "Deleting /home/chow/Softs/mysql"
    sudo rm -rf /home/chow/Softs/mysql
fi

if [ -d "/usr/local/mysql" ];then
    echo "Deleting /usr/local/mysql"
    sudo rm -rf /usr/local/mysql
fi

# delete cnf file
if [ -f "/etc/my.cnf" ];then
    echo "Deleting /etc/my.cnf"
    sudo rm /etc/my.cnf
fi

# delete user/group
sudo userdel mysql
sudo groupdel mysql

echo "============================"