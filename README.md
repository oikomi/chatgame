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

重名
<pre><code>
root@ubuntu:/# telnet 127.0.0.1 8080
Trying 127.0.0.1...
Connected to 127.0.0.1.
Escape character is '^]'.
Please input your username :
/setName zhangsan
username is already used, Please input your username :
</code></pre>


60超时
======
在client结构中加一个Alive字段，起一个定时器60去遍历clients, Alive为false,则干掉连接。 注意：在read里读到数据会把Alive设为true

