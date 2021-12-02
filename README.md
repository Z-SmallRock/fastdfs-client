根据开源项目对配置文件的使用进行了简单修改
原有开源项目是有track.conf 如果我们自己有自己的配置文件，导致配置文件过多。
目前修改成一个yaml里面。统一配置
# fdfs_client.go

fastdfs go client implement

**1 support** 

upload(UploadByFilename,UploadByBuffer)

download(DownloadToFile,DownloadToBuffer)

delete(DeleteFile)

**2 append is not support(limited by fastdfs server)**

you can implement append with delete origin && upload new again

**3 UploadByFilename realized with sendfile syscall in linux,so UploadByBuffer is depracated**

**4 realized conn_pool,pool_size control by config file**

**5 details see client_test.go,good luck ^_^**

## $ go get git@github.com:Z-SmallRock/fastdfs-client.git

