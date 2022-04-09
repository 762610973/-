#  Git

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

`git log --graph --pretty=format:"%h %s"`

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

==遇到冲突需要手动解决，并且`git commit -m "description"`后面不加文件名（冲突原因是两个分支在同一个文件的同一个位置有两套完全不同的修改）==

### git rebase,这一块儿有待加强

==git rebase==使提交记录变得更简洁

1. 将多条记录合并为一条，让自己的提交记录变得更简单
   合并记录时，建议不要合并已经push到仓库的

   `git rebase -i  HEAD~n`

   `git rebase -i versionNode`

2. 消除分支
   `git rebase main //分支上操作这个`
   `git merge branchName //主干上操作这个`

3. 分叉
   `git fetch`

   `git rebase`



### 版本回滚

`git reset --hard HEAD //回退到上个版本`

`git reset --hard HEAD~n //回退到n次提交之前`

`git reset --hard commit _id 回退快进到指定的commit的哈希码`

`git reset HEAD file //取消暂存file`

`git checkout--file // 撤销之前的修改`

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

- `git push 别名/链接 分支`
- `git push <远程主机名> <本地分支名>:<远程分支名>`
- `git push <远程主机名> <本地分支名> //本地分支名与远程分支名相同，省略冒号`
- `git push origin //将当前分支推送到origin主机的对应分支`
- ` git push -u origin master //此命令将本地的master分支推送到origin主机，同时指定origin为默认主机，后面就可以不加任何参数使用git push了。`
- `git add. || git commit -m "说明" || git remote add origin 仓库链接地址 || git pull origin master// 首次提交要git pull 一下 || git push -u origin main`

#### 克隆远程仓库到本地

- `git clone 链接 //公共库放开读权限`
- ==clone会做如下操作：1、拉去代码。2、初始化本地仓库。3、创建别名。==

#### 邀请加入团队

- github操作

#### 拉取远程库内容

- `git pull 远程仓库地址别名 远程分支名`

#### 跨团队

`Fork + create pull request + Pull request`

#### SSH免密登录

`ssh-keygen -t rsa -C email`

# Gitee码云

==Gitee与Git极为类似==

==在Goland中的操作也一样==

