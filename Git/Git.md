# Git

## Git概述

- 免费、开源、分布式版本控制系统

- 性能优于Subversion、CVS、Perforce和ClearCase

- 版本控制

  > 集中式：单一服务器，每个开发者能看到其他开发者的进度等，管理员可以轻松掌控每个   		          开发者的权限，缺点是中央服务器的单点故障
  >
  > 分布式：无单点故障，每个客户端都是整个完整的项目

- Git：由C语言编写，Linus自己用C语言开发的分布式版本控制系统，主体程序开发完成只用了两周
- - 工作区：写代码（git add到暂存区）
  - 暂存区：临时存储（git commit到本地库，生成历史版本）
  - 本地库：历史版本

## Git常用命令

### 基本语法

`git 按行来管理`

`git config --global init.defaultBranch main //将Git默认分支从master修改为main`

`git config --global user.name 用户名 //设置用户签名，只是本地的`

`git config --global user.email 邮箱 //设置用户签名，只是本地的`
`git init //初始化本地库`

`git status //查看本地库状态`

`git add file //添加文件到暂存区`

`git rm --cached file //删除暂存区的文件 `

`git commit -m "日志信息" file //将暂存区的文件添加到本地库`

`git reflog //查看版本号，前7位`

`git log //查看版本信息（完整信息包括：完整的版本号和提交的时间以及提交作者）`

### 穿梭版本

`git reflog --hard 版本号 //穿梭到版本号对应的版本`

==git切换分支是通过HEAD指针==

==head指向master，master指向版本，需要穿梭版本的时候，只需改变master的指向即可==

### Git分支操作

==同时并行推进多个功能开发，提高开发效率==

`git branch -v //查看分支 `==代表当前所在的分支==

`git branch name //创建分支`

`git branch //查看本地所有分支`

`git branch -r //查看远程所有分支`

`git branch -a //查看本地和远程所有分支`

`git branch -f branchName //新建一个分支但不切换`

`git branch -d branchName //删除本地分支`

`git branch -d -r branchName //删除远程分支，删除后需要推动到服务器，git push origin :branchName,-D表示强制删除`

`git branch -m oldbranch newbranch //重命名分支，-M强制重命名，如果重命名分支需要：1.删除远程修改分支 2.push本地新分支名到远程`

`git checkout -b branchName //新建并切换新分支`

`git checkout branchName //切换分支`

`git merge branchName //将branchName合并到当前所在的分支`

==遇到冲突需要手动解决，并且`git commit -m "description"`后面不加文件名==

### 团队协作与跨团队协作

# Github

### 创建远程仓库

- 远程库和本地库名称最好**相同**

### 远程仓库操作

#### 创建远程仓库别名

- `git remote -v //查看当前有什么别名`
- `git remote add 别名 链接地址`
- ![image-20211228202306594](https://picgo-for-typora.oss-cn-beijing.aliyuncs.com/image-20211228202306594.png)

#### 推送本地分支到远程仓库

- `git push 别名 分支`

#### 克隆远程仓库到本地

- 

#### 邀请加入团队

- 

#### 拉去远程库内容





# Gitee码云



# GitLab



