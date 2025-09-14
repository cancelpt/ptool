# ptool

功能介绍请参考源仓库 https://github.com/sagan/ptool

# 改动

## 批量下载种子 (batchdl)

提供一个 batchdl 命令用于批量下载 PT 网站的种子。默认按种子体积大小升序排序、跳过死种和已经下载过的种子。

**该版本增加**：

- --no-useless : 跳过做种人数大于 30 且做种人数大于完成人数 3 倍以上的种子（已将你的做种计算在内）。
- --no-free : 跳过免费的种子。
- --published-before : 跳过一段时间发布的新种子，单位为分钟，例如跳过 48 小时内发布的种子 `--published-before 2880`。

### ptool 的一些实际使用场景示例：

#### 1. 下载非隐性中性且限定体积和做种人数的种子

```
ptool batchdl 站点名称 --sort seeders --min-torrent-size 3.2GiB --max-torrent-size 5.9GiB --min-seeders 1 --max-seeders 3 --no-useless --no-neutral --base-url "torrents.php?p1=1" --download --slow --skip-existing --download-dir D:\\ptool\\站点名称_torrents -vv --timeout 10 --published-before 2880
```

**解析** ：

- 使用 `seeders`作为`--sort`，默认做种人数多的在前面，但`ptool`总是从最后一页开始抓取（是为了排除置顶种子的影响），等同于从做种人数少的种子开始抓取；
- `--min-torrent-size`和`--max-torrent-size`框定了种子的最大体积和最小体积；
- `--min-seeders`和`--max-seeders`框定了种子的最大做种人数和最小做种人数；
- `--no-useless`跳过了隐性中性，`--no-neutral`跳过了标签中含有中性的种子；
- `--base-url`可以自行配置，例如在站点搜索箱中过滤官种等等，然后搜索一次，将其`url`作为`base-url`，需要注意的是`ptool`的种子列表排序由程序自动补齐，不要在`base-url`中带有排序的参数（换言之，不要传入一个类似按种子体积排序的`base-url`）；
- `--download`下载种子而不是推送到下载器，因为一些站点有流控，例如每小时 100 个或者每 4 小时 200 个等等，如果使用推送到下载器，可能造成种子未能在下载器中下载；`--slow`确保种子慢慢下载，不进行高额并发，`--skip-existing`跳过下载目录`--download-dir`里面有的种子；
- `-vv`便于调试；
- `--timeout 10`，对于访问不太顺利的站点，允许更长的超时时间；
- `--published-before 2880`跳过 48 小时内发布的种子。

**额外注意！！！**

使用上述命令的时候，你应该要注意`ptool`因为站点流控无法下载种子停止时，种子页在哪一页

此时`ptool`会有类似如下输出：

```
torrent size.12345 (Some Torrent Name): failed to download: server return invalid content-type: text/html
```

最后一行会有类似`LastPage:`的页号，记录下来，等流控过后，使用`--start-page`传入该页号，避免频繁访问同一个种子列表页面

假设上次停止的页号是`3203`，则传入`--start-page 3203`：

```
ptool batchdl 站点名称 --sort seeders --min-torrent-size 3.2GiB --max-torrent-size 5.9GiB --min-seeders 1 --max-seeders 3 --no-useless --no-neutral --base-url "torrents.php?p1=1" --download --slow --skip-existing --download-dir D:\\ptool\\站点名称_torrents -vv --timeout 10 --start-page 3203
```

#### 2. 获取种子列表形成`json`文件

```
ptool batchdl 站点名称 --base-url "torrents.php?cat%5B%5D=412&cat%5B%5D=405" --save-json-file 站点名称.json --exclude DV,DoVi,HDR,EDR,fps --sort time --order desc
```

**解析** ：

- 使用种子发布时间`time`作为`--sort`，`--order`为`desc`降序，这样可以让`ptool`优先获取新种子；
- `--save-json-file`给出了保存`json`文件的路径；
- `--exclude`可以用来排除一些你不想要的种子，例如设备无法播放杜比视界 HDR 等；

`ptool`会将种子信息保存成类似如下的`json`文件，用于你后续的工作流：

```
    {
        "Name": "Some Name S01 2025 2160p WEB-DL H265 AAC-HELLO",
        "Description": "全100集 | 类型: 剧情 | 主演: 张三 / 李四",
        "Id": "ptsite.12345",
        "InfoHash": "",
        "DownloadUrl": "https://pt.pt/download.php?id=12345\u0026letdown=1",
        "DownloadMultiplier": 0,
        "UploadMultiplier": 1,
        "DiscountEndTime": 1712312312,
        "Time": 1712312312,
        "Size": 123456789,
        "IsSizeAccurate": false,
        "Seeders": 3,
        "Leechers": 0,
        "Snatched": 23,
        "HasHnR": false,
        "IsActive": false,
        "IsCurrentActive": false,
        "Paid": false,
        "Bought": false,
        "Neutral": false,
        "Tags": [
            "TV-Pack",
            "WEB-DL"
        ]
    },
```
