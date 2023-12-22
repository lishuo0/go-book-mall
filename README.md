//搜索nginx镜像
docker search nginx
//下载nginx镜像（下载最新版本latest）
docker pull nginx

//容器里Nginx默认配置文件在/etc/nginx目录
//-v目录映射，本地主机文件与容器文件建立映射关系
//-p 端口映射
//-d 后台模式运行，输出容器ID
docker run --name nginx -p 80:80 \
-v /xxx/nginx/nginx.conf:/etc/nginx/nginx.conf \
-v /xxx/nginx/conf.d:/etc/nginx/conf.d \
-v /xxx/nginx/logs:/var/log/nginx \
-d nginx:latest



//搜索Mysql镜像
docker search mysql
//下载8.0版本Mysql镜像
docker pull mysql:8.0
// 启动Mysql实例
docker run --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d -p 3306:3306 mysql:8.0

