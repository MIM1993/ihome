#启动redis
sudo redis-server /etc/redis/redis.conf

#启动fastdfs
sudo  fdfs_trackerd   /home/itcast/workspace/go/src/sss/IhomeWeb/conf/tracker.conf  restart
sudo   fdfs_storaged  /home/itcast/workspace/go/src/sss/IhomeWeb/conf/storage.conf  restart

#启动nginx
sudo nginx