# 安装
## x86芯片的linux
- 解压: unzip msm-linux-amd64.zip
- 进入目录: cd msm-linux-amd64
## arm芯片的linux
- 解压: unzip msm-linux-arm64.zip
- 进入目录: cd msm-linux-arm64


# 配置
- 编辑config/msm.yaml
    - 将source.file的值设置为建表语句文件的路径
    - 将target.file的值设置为建表语句文件的路径

# 显示表结构差异
- 在linux shell里运行./show-diff.sh
- 输出内容为json格式, 表示表, 字段, 索引的差异

# 显示表结构迁移脚本
- 在linux shell里运行./show-migration.sh
- 输出内容为sql语句列表, 其中DROP TABLE相关语句被注释
- 对于DROP TABLE或DROP COLUMN等高危语句, 请确认无问题后谨慎执行
