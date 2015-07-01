编译安装
进入主目录  make

运行

./bin/server config/server.json​

客户端测试
<pre><code>
root@ubuntu:/# telnet 127.0.0.1 8080
Trying 127.0.0.1...
Connected to 127.0.0.1.
Escape character is '^]'.
Please input your username :
/setName zhangsan
Notification:  changed name to zhangsan
Welcome  zhangsan
 to join the ChatRoom
Current have 1 people on line
/weather beijing
beijing
 is sunnshine
/ticket beijing shanghai
from beijing to shanghai
 is sell out
/quit
Connection closed by foreign host.
</code></pre>
