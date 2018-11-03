# AgendaCLI

## Usage

- **用户注册**：`AgendaCLI register -u=用户名 -p=密码 -e=邮箱 -t=电话`

![1541227525599](/images/1541227924601.png)

- **用户登录**：`AgendaCLI login -u=用户名 -p=密码`

![1541227976941](/images/1541227976941.png)

- **用户查询**：`AgendaCLI searchUser`

![1541228004819](/images/1541228004819.png)

- **创建会议**：`AgendaCLI createMeeting -t=会议主题 -p=会议参与者 -s=开始时间 -e=结束时间`


![1541228129916](/images/1541228129916.png)

- **增删会议参与者**：

	**增**：`AgendaCLI operateParticipant -t=会议主题 -o=add -p=参与者`

	**删**:   `AgendaCLI operateParticipant -t=会议主题 -o=del -p=参与者`

![1541228190175](/images/1541228190175.png)

- **查询会议**：`AgendaCLI searchMeeting -s=查询的起始时间 -e=查询的结束时间`

![1541228216982](/images/1541228216982.png)

- **取消会议**：`AgendaCLI cancelMeeting -t=会议主题`

![1541228238740](/images/1541228238740.png)

- **退出会议**：`AgendaCLI exitMeeting -t=会议主题`

![1541228262853](/images/1541228262853.png)

- **清空会议**:  `AgendaCLI deleteAllMeeting`

![1541228320161](/images/1541228320161.png)

- **删除用户**:  `AgendaCLI deleteUser`

![1541228338809](/images/1541228338809.png)

- 创建的用户保存到`User.json`文件


![1541228426473](/images/1541228426473.png)

- 创建的会议保存到`Meeting.json`文件

![1541228486557](/images/1541228486557.png)

