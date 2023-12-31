package main

import (
	"fmt"
	"io"

	"github.com/cloudfoundry/jibber_jabber"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	initEnglish()
	initChinese()
}

func initEnglish() {
	// main.go
	message.SetString(language.English, "parse Option error", "Parse Option error")
	message.SetString(language.English, "couldn't find Git execute program: %s", "Couldn't find Git execute program: %s")
	message.SetString(language.English, "sorry, this tool requires Git version at least 2.24.0",
		"Sorry, this tool requires Git version at least 2.24.0")
	message.SetString(language.English, "couldn't support running in bare repository", "Couldn't support running in bare repository")
	message.SetString(language.English, "couldn't support running in shallow repository", "Couldn't support running in shallow repository")
	message.SetString(language.English, "scanning repository error: %s", "Scanning repository error: %s")
	message.SetString(language.English, "no files were scanned",
		"According to the filter conditions you selected, no files were filtered out. Please adjust the filter criteria and try again.")
	message.SetString(language.English, "no files were selected",
		"You haven't selected any files. Please select at least one file")
	message.SetString(language.English, "operation aborted", "The operation has been aborted. Please reconfirm the file and try again.")
	message.SetString(language.English, "cleaning completed", "Local repository cleaning up completed!")
	message.SetString(language.English, "current repository size", "Current repository size: ")
	message.SetString(language.English, "including LFS objects size", "LFS objects size: ")
	message.SetString(language.English, "execute force push",
		"The following two commands will be executed, then the remote commit will be overwritten:")
	message.SetString(language.English, "suggest operations header",
		"Finally, please confirm the current repo status is Ok and no file is deleted by mistake and the repo size is under the repo size limit, please follow those steps below:")
	message.SetString(language.English, "1. (Done!)", "1. (Done!) remote repository have been updated.")
	message.SetString(language.English, "1. (Undo)", "1. (Undo) update remote repository. Push local cleaned repository to remote repository:")
	message.SetString(language.English, "2. (Undo)",
		"2. (Undo) clean up the remote repository. After successful push, please go to your corresponding repository management page to perform GC operation.")
	message.SetString(language.English, "3. (Undo)",
		"3. (Undo) process the associated repository. Process other repository in the clone under the same remote repository to ensure that the same file won't be submitted to the remote repository again. ")
	message.SetString(language.English, "gitee GC page link", "    please click Gitee repo manage link: ")
	message.SetString(language.English, "for detailed documentation, see", "    For detailed documentation, see: ")
	message.SetString(language.English, "introduce GIT LFS",
		"If you have Gitee LFS(large file storage) service,  you can use '--lfs' option to convert big file into LFS to manage your large file separately.")
	message.SetString(language.English, "for the use of Gitee LFS, see", "For the use of Gitee LFS, see: ")
	message.SetString(language.English, "init repo filter error", "Init repo Filter error")
	message.SetString(language.English, "ask question module fail: %s", "Ask question module fail: %s")
	message.SetString(language.English, "before you push to remote, you have to do something below:",
		"Before you push to remote, you have to do something below:")
	message.SetString(language.English, "1. install git-lfs",
		"1. install git-lfs by this link: https://packagecloud.io/github/git-lfs/install")
	message.SetString(language.English, "2. run command: git lfs install",
		"2. run command: git lfs install")
	message.SetString(language.English, "3. edit .gitattributes file",
		`3. run command: git lfs track "your-file"(this will modify .gitattributes file`)
	message.SetString(language.English, "4. commit your .gitattributes file.",
		"4. commit your .gitattributes file.")

	// options.go
	message.SetString(language.English, "help info", Usage)
	message.SetString(language.English, "option format error: %s", "Option format error: %s")
	message.SetString(language.English, "build version: %s", "Build version: %s")
	message.SetString(language.English, "single parameter is invalid", "This single parameter is invalid, please combine with other parameter.")
	message.SetString(language.English, "LFS parameter is invalid", "--lfs parameter must combine with --scan and --type parameter.")
	// parser.go
	message.SetString(language.English, "unsupported filechange type", "Unsupported filechange type")
	message.SetString(language.English, "nested tags error",
		"The operation has been aborted because nested tags. It is recommended to use the '--branch=<branch>' option to specify a single branch.")
	message.SetString(language.English, "no match mark id", "No match mark id")
	message.SetString(language.English, "no match original-oid", "No match original-oid")
	message.SetString(language.English, "no match data size", "No match data size")
	message.SetString(language.English, "failed to write data", "Failed to write data")
	message.SetString(language.English, "start to clean up specified files",
		"Start to clean up the specified file from the history (if the repository is too large, the execution time will be long, please wait a few minutes)...")
	message.SetString(language.English, "start to migrate specified files",
		"Start converting the specified file to an LFS file (if the repository is too large, the execution time will be long, please wait a few minutes)...")

	message.SetString(language.English, "run git-fast-import process failed", "Run git-fast-import process failed")
	// utils.go
	message.SetString(language.English, "expected a value followed by --limit option, but you are: %s",
		"Expected a value followed by --limit option, but you are: %s")
	message.SetString(language.English, "expected format: --limit=<n>b|k|m|g, but you are: --limit=%s",
		"Expected format: --limit=<n>b|k|m|g, but you are: --limit=%s")
	message.SetString(language.English, "scan done!", "Scan done!")
	message.SetString(language.English, "note that there may be multiple versions of the same file",
		"Note that there may be multiple versions of the same file, which are the main reasons for wasting git repository storage")

	// repository.go
	message.SetString(language.English, "start scanning", "Start scanning(if the repository is too large, the scanning time will be long, please wait a few minutes)...")
	message.SetString(language.English, "run GetBlobName error: %s", "Run GetBlobName error: %s")
	message.SetString(language.English, "run getblobsize error: %s", "Run getblobsize error: %s")
	message.SetString(language.English, "expected blob object type, but got: %s", "Expected blob object type, but got: %s")
	message.SetString(language.English, "could not run 'git rev-parse --is-bare-repository': %s", "Could not run 'git rev-parse --is-bare-repository': %s")
	message.SetString(language.English, "could not run 'git rev-parse --is-shallow-repository': %s", "Could not run 'git rev-parse --is-shallow-repository': %s")
	message.SetString(language.English, "could not run 'git reflog show': %s", "Could not run 'git reflog show': %s")
	message.SetString(language.English, "could not run 'git lfs version': %s", "Could not run 'git lfs version': %s")
	message.SetString(language.English, "could not run 'git version': %s", "Could not run 'git version': %s")
	message.SetString(language.English, "match git version wrong", "Match git version wrong")
	message.SetString(language.English, "could not run 'git symbolic-ref HEAD --short': %s", "Could not run 'git symbolic-ref HEAD --short': %s")
	message.SetString(language.English, "could not run 'git status'", "Could not run 'git status', make sure you are in a git repo.")
	message.SetString(language.English, "there's some changes to be committed, please commit them first",
		"There's some changes to be committed, please commit them first(Try to use 'git status' to see un-committed changes).")
	message.SetString(language.English, "could not run 'du -hs'", "Could not run 'du -hs'")
	message.SetString(language.English, "backup canceled", "Backup canceled")
	message.SetString(language.English, "start backup", "Start backup...")
	message.SetString(language.English, "clone error", "git clone --no-local error")
	message.SetString(language.English, "run filepach.Abs error", "Run filepach.Abs error")
	message.SetString(language.English, "bare repo warning", "⚠ Warning: you are in a bare or mirror repo, some operations may be limited.")
	message.SetString(language.English, "bare repo error", "❌ Error: can't perform any LFS operation in a bare or mirror repo.")

	message.SetString(language.English, "backup done! Backup file path is: %s", "Backup done! Backup file path is: %s")
	message.SetString(language.English, "push failed",
		"Push failed. Your repo may still exceed the size limit, please clear other history big files again, and then push by hand.")
	message.SetString(language.English, "done", "Done")
	message.SetString(language.English, "file cleanup is complete. Start cleaning the repository", "File cleanup is complete. Start cleaning the repository...")
	message.SetString(language.English, "branches have been changed", "The following branches have been changed: ")
	message.SetString(language.English, "nothing have changed, exit...", "Nothing have changed, exit...")

	// cmd.go
	message.SetString(language.English, "select the type of file to scan, such as zip, png:", "Select the type of file to scan, such as zip, png:")
	message.SetString(language.English, "default is all types. If you want to specify a type, you can directly enter the type suffix without prefix '.'",
		"Default is all types. If you want to specify a type, you can directly enter the type suffix without prefix '.'")
	message.SetString(language.English, "filetype error one", "Sorry, the type name you entered is too long, more than 50 characters")
	message.SetString(language.English, "filetype error two", "The type must be a letter. It can contain '.' in the middle, but it doesn't need to contain '.' at the beginning")

	message.SetString(language.English, "select the minimum size of the file to scan, such as 1m, 1G:", "Select the minimum size of the file to scan, such as 1m, 1G:")
	message.SetString(language.English, "the size value needs units, such as 10K. The optional units are B, K, m and G, and are not case sensitive",
		"The size value needs units, such as 10K. The optional units are B, K, m and G, and are not case sensitive")
	message.SetString(language.English, "filesize error one", "input error")
	message.SetString(language.English, "filesize error two", "Must be a combination of numbers + unit characters (B, K, m, g), and the units are not case sensitive")

	message.SetString(language.English, "select the number of scan results to display, the default is 3:", "Select the number of scan results to display. The default value is 3:")
	message.SetString(language.English, "the default display is the first 3. The maximum page size is 10 rows, so it is best not to exceed 10.",
		"The default display is the first 3. The maximum page size is 10 rows, so it is best not to exceed 10.")
	message.SetString(language.English, "filenumber error one", "input error")
	message.SetString(language.English, "filenumber error two", "Must be a pure number")

	message.SetString(language.English, "multi select message", "Please select the file you want to delete (multiple choices are allowed):")
	message.SetString(language.English, "multi select help info", "Use <Up/Down> arrows to move, <space> to select, <right> to all, <left> to none, type to filter, ? for more help")

	message.SetString(language.English, "confirm message", "The above is the file you want to delete. Are you sure you want to *DELETE* it ?")
	message.SetString(language.English, "ask for override message", "A folder with the same name exists in the current directory. Please make sure you do not overwrite the original backup file")
	message.SetString(language.English, "ask for update message", "You have done a repo clear work. You can force push to remote if no file is deleted by mistake and the repo size is under the limit. Otherwise, select NO to continue clearing history files.")
	message.SetString(language.English, "ask for migrating big file into LFS", "Do you want to migrate your big files into Gitee LFS? ")
	message.SetString(language.English, "process interrupted", "process interrupted")

	message.SetString(language.English, "convert uint error: %s", "Convert uint error: %s")
	message.SetString(language.English, "parse uint error: %s", "Parse uint error: %s")

	message.SetString(language.English, "file have been changed", "those files have been converted to LFS file:")
	message.SetString(language.English, "Convert LFS file error", "Can not convert noraml file into LFS file under non-scan mode")

}

func initChinese() {
	// main.go
	message.SetString(language.Chinese, "parse Option error", "解析参数错误。")
	message.SetString(language.Chinese, "couldn't find Git execute program: %s", "无法找到Git可执行文件: %s")
	message.SetString(language.Chinese, "sorry, this tool requires Git version at least 2.24.0",
		"抱歉，这个工具需要Git的最低版本为 2.24.0")
	message.SetString(language.Chinese, "couldn't support running in bare repository", "不支持在裸仓库中执行。")
	message.SetString(language.Chinese, "couldn't support running in shallow repository", "不支持在浅仓库中执行。")
	message.SetString(language.Chinese, "scanning repository error: %s", "扫描仓库出错: %s")
	message.SetString(language.Chinese, "no files were scanned", "根据你所选择的筛选条件，没有扫描到任何文件，请调整筛选条件再试一次。")
	message.SetString(language.Chinese, "no files were selected", "你没有选择任何文件，请至少选择一个文件。")
	message.SetString(language.Chinese, "operation aborted", "操作已中止，请重新确认文件后再次尝试。")
	message.SetString(language.Chinese, "cleaning completed", "本地仓库清理完成！")
	message.SetString(language.Chinese, "current repository size", "当前仓库大小：")
	message.SetString(language.Chinese, "including LFS objects size", "LFS对象大小:  ")
	message.SetString(language.Chinese, "execute force push", "将会执行如下两条命令，远端的的提交将会被覆盖:")
	message.SetString(language.Chinese, "suggest operations header", "最后，当确认当前仓库状态是正常，不存在文件误删除，并且未超过仓库大小限制后，请手动完成如下工作：")
	message.SetString(language.Chinese, "1. (Done!)", "1. (已完成！)远程仓库已经更新。")
	message.SetString(language.Chinese, "1. (Undo)", "1. (待完成)更新远程仓库。将本地清理后的仓库手动推送到远程仓库：")
	message.SetString(language.Chinese, "2. (Undo)", "2. (待完成)清理远程仓库。提交成功后，请前往你对应的仓库管理页面，执行GC操作。")
	message.SetString(language.Chinese, "3. (Undo)", "3. (待完成)处理关联仓库。处理同一远程仓库下clone的其它仓库，确保不会将同样的文件再次提交到远程仓库。")
	message.SetString(language.Chinese, "gitee GC page link", "    请点击Gitee仓库管理页面链接: ")
	message.SetString(language.Chinese, "for detailed documentation, see", "    详细文档请参阅: ")
	message.SetString(language.Chinese, "introduce GIT LFS", "如果开通了Gitee LFS(Large file storage)服务，可使用'--lfs'选项，将大文件迁移到LFS服务器进行管理。")
	message.SetString(language.Chinese, "for the use of Gitee LFS, see", "Gitee LFS 的使用请参阅：")
	message.SetString(language.Chinese, "init repo filter error", "初始化仓库过滤器失败")
	message.SetString(language.Chinese, "ask question module fail: %s", "交互式模块运行失败: %s")
	message.SetString(language.Chinese, "before you push to remote, you have to do something below:",
		"在你推送仓库到远程之前，必须完成以下操作：")
	message.SetString(language.Chinese, "1. install git-lfs",
		"1. 安装 Git LFS: https://packagecloud.io/github/git-lfs/install")
	message.SetString(language.Chinese, "2. run command: git lfs install",
		"2. 在仓库中运行命令：git lfs install ")
	message.SetString(language.Chinese, "3. edit .gitattributes file",
		`3. 追踪上述文件(这会修改 .gitattributes 文件)：git lfs track "your-file"`)
	message.SetString(language.Chinese, "4. commit your .gitattributes file.",
		"4. 提交修改后的 .gitattributes 文件")

	// options.go
	message.SetString(language.Chinese, "help info", Usage_ZH)
	message.SetString(language.Chinese, "option format error: %s", "选项格式错误: %s")
	message.SetString(language.Chinese, "build version: %s", "版本编号: %s")
	message.SetString(language.Chinese, "single parameter is invalid", "该单项参数无效，请结合其它参数使用")
	message.SetString(language.Chinese, "LFS parameter is invalid", "--lfs 选项必须结合选项 --scan 和选项 --type 使用")
	// parser.go
	message.SetString(language.Chinese, "unsupported filechange type", "不支持的filechange类型")
	message.SetString(language.Chinese, "nested tags error", "处理过程中断，因为仓库中存在嵌套式tag，建议使用'--branch=<branch>'参数指定单个分支。")
	message.SetString(language.Chinese, "no match mark id", "没有匹配到mark id字段")
	message.SetString(language.Chinese, "no match original-oid", "没有匹配到original oid字段")
	message.SetString(language.Chinese, "no match data size", "没有匹配到数据大小字段")
	message.SetString(language.Chinese, "failed to write data", "写数据失败")
	message.SetString(language.Chinese, "start to clean up specified files", "开始从历史中清理指定的文件(如果仓库过大，执行时间会比较长，请耐心等待)...")
	message.SetString(language.Chinese, "start to migrate specified files", "开始将指定的文件转换为LFS文件(如果仓库过大，执行时间会比较长，请耐心等待)...")
	message.SetString(language.Chinese, "run git-fast-import process failed", "运行git fast-import过程出错")
	// utils.go
	message.SetString(language.Chinese, "expected a value followed by --limit option, but you are: %s", "'--limit'选项后面需要跟一个数值，但是你给的是: %s")
	message.SetString(language.Chinese, "expected format: --limit=<n>b|k|m|g, but you are: --limit=%s", "希望的格式为: --limit=<n>b|k|m|g, 但是你给的是: --limit=%s")
	message.SetString(language.Chinese, "scan done!", "扫描完成!")
	message.SetString(language.Chinese, "note that there may be multiple versions of the same file", "注意，同一个文件因为版本不同可能会存在多个。")
	// repository.go
	message.SetString(language.Chinese, "start scanning", "开始扫描(如果仓库过大，扫描时间会比较长，请耐心等待)...")
	message.SetString(language.Chinese, "run GetBlobName error: %s", "运行 GetBlobName 错误: %s")
	message.SetString(language.Chinese, "run getblobsize error: %s", "运行 getblobsize 错误: %s")
	message.SetString(language.Chinese, "expected blob object type, but got: %s", "期望blob类型数据，但实际得到: %s")
	message.SetString(language.Chinese, "could not run 'git rev-parse --is-bare-repository': %s", "无法运行'git rev-parse --is-bare-repository': %s")
	message.SetString(language.Chinese, "could not run 'git rev-parse --is-shallow-repository': %s", "无法运行'git rev-parse --is-shallow-repository': %s")
	message.SetString(language.Chinese, "could not run 'git reflog show': %s", "无法运行'git reflog show': %s")
	message.SetString(language.Chinese, "could not run 'git lfs version': %s", "无法运行'git lfs version': %s")
	message.SetString(language.Chinese, "could not run 'git version': %s", "无法运行'git version': %s")
	message.SetString(language.Chinese, "match git version wrong", "Git版本号匹配错误")
	message.SetString(language.Chinese, "could not run 'git symbolic-ref HEAD --short': %s", "无法运行'git symbolic-ref HEAD --short': %s")
	message.SetString(language.Chinese, "could not run 'git status'", "无法运行`git status`, 请确保你在一个 Git 仓库中")
	message.SetString(language.Chinese, "there's some changes to be committed, please commit them first", "当前仍有未提交的更改，请先提交(使用 git status 查看未提交的更改)。")
	message.SetString(language.Chinese, "git status clean", "git status为空")
	message.SetString(language.Chinese, "could not run 'du -hs'", "无法运行'du -hs'")
	message.SetString(language.Chinese, "backup canceled", "已取消备份")
	message.SetString(language.Chinese, "start backup", "开始备份...")
	message.SetString(language.Chinese, "clone error", "git clone --no-local 错误")
	message.SetString(language.Chinese, "run filepach.Abs error", "运行 filepach.Abs 错误")
	message.SetString(language.Chinese, "backup done! Backup file path is: %s", "备份完毕! 备份文件路径为：%s")
	message.SetString(language.Chinese, "Push failed", "推送失败，可能是仓库大小仍然超出限制，建议继续清理其它历史大文件，再手动推送。")
	message.SetString(language.Chinese, "done", "完成")
	message.SetString(language.Chinese, "file cleanup is complete. Start cleaning the repository", "文件清理完毕，开始清理仓库...")
	message.SetString(language.Chinese, "branches have been changed", "以下分支已经更改：")
	message.SetString(language.Chinese, "nothing have changed, exit...", "没有文件更改，退出...")
	message.SetString(language.Chinese, "bare repo warning", "⚠ 警告：您正在裸仓或镜像仓中，有些操作可能会受到限制。")
	message.SetString(language.Chinese, "bare repo error", "❌ 错误：无法在裸仓或镜像仓中执行LFS相关操作!")

	// cmd.go
	message.SetString(language.Chinese, "select the type of file to scan, such as zip, png:", "选择要扫描的文件的类型，如：zip, png:")
	message.SetString(language.Chinese, "filetype help message", "默认无类型，即查找所有类型。如果想指定类型，则直接输入类型后缀名即可, 不需要加'.'")
	message.SetString(language.Chinese, "filetype error one", "抱歉，输入的类型名过长，超过50个字符")
	message.SetString(language.Chinese, "filetype error two", "类型必须是字母，中间可以包含'.'，但是开头不需要包含'.'")

	message.SetString(language.Chinese, "select the minimum size of the file to scan, such as 1m, 1G:", "选择要扫描文件的最低大小，如：1M, 1g:")
	message.SetString(language.Chinese, "filesize help message", "大小数值需要单位，如: 10K. 可选单位有B,K,M,G, 且不区分大小写")
	message.SetString(language.Chinese, "filesize error one", "输入错误")
	message.SetString(language.Chinese, "filesize error two", "必须以数字+单位字符(b,k,m,g)组合，且单位不区分大小写")

	message.SetString(language.Chinese, "select the number of scan results to display, the default is 3:", "选择要显示扫描结果的数量，默认值是3:")
	message.SetString(language.Chinese, "filenumber help message", "默认显示前3个，单页最大显示为10行，所以最好不超过10。")
	message.SetString(language.Chinese, "filenumber error one", "输入错误")
	message.SetString(language.Chinese, "filenumber error two", "必须是纯数字")

	message.SetString(language.Chinese, "multi select message", "请选择你要删除的文件(可多选):")
	message.SetString(language.Chinese, "multi select help info", "使用键盘的上下左右，可进行上下换行、全选、全取消，使用空格建选中单个，使用Enter键确认选择。")

	message.SetString(language.Chinese, "confirm message", "以上是你要删除的文件，确定要<删除>吗?")
	message.SetString(language.Chinese, "ask for override message", "检测到存在同名备份文件，你可能已经备份，请谨慎操作，不要覆盖原始备份！选择No/N取消备份。")
	message.SetString(language.Chinese, "ask for update message", "你已完成一次文件清理过程，请确认仓库没有文件误删除，且仓库大小已经满足推送条件，则可以强制推送，否则选择No, 继续清理其它文件。")
	message.SetString(language.Chinese, "ask for migrating big file into LFS", "是否将大文件迁移到 Gitee LFS 进行管理？")
	message.SetString(language.Chinese, "process interrupted", "过程中断")

	message.SetString(language.Chinese, "convert uint error: %s", "转换大小单位出错: %s")
	message.SetString(language.Chinese, "parse uint error: %s", "解析无符号整数出错: %s")

	message.SetString(language.Chinese, "file have been changed", "以下这些文件已经被转化为 LFS 文件:")
	message.SetString(language.Chinese, "Convert LFS file error", "无法在非扫描模式下将文件转换为 LFS 文件")

}

// find local languange type. LC_ALL > LANG > LANGUAGE
func Local() language.Tag {
	// when LC_ALL && LANG is none, throw panic
	userLanguage, err := jibber_jabber.DetectLanguage()
	if err != nil {
		fmt.Println("try to set: export LC_ALL=zh_CN.UTF-8")
		panic(err)
	}
	// fix LC_ALL=C.UTF-8
	if userLanguage == "C" {
		userLanguage = "zh"
	}
	tagLanguage := language.Make(userLanguage)
	return tagLanguage
}

// set languange
func SetLang() language.Tag {
	return Local()
}

// local printer
func LocalPrinter() *message.Printer {
	tag := SetLang()
	return message.NewPrinter(tag)
}

// local fmt.Sprintf
func LocalSprintf(key message.Reference, a ...interface{}) string {
	return LocalPrinter().Sprintf(key, a)
}

// local fmt.Fprintf
func LocalFprintf(w io.Writer, key message.Reference, a ...interface{}) (int, error) {
	return LocalPrinter().Fprintf(w, key, a)
}

// local fmt.Printf
func LocalPrintf(key message.Reference, a ...interface{}) (int, error) {
	return LocalPrinter().Printf(key, a)
}

/*  PRINT WITH COLOR */

func PrintLocalWithRed(key message.Reference, a ...interface{}) {
	f := LocalSprintf(key, a)
	PrintRed(f)
}

func PrintLocalWithGreen(key message.Reference, a ...interface{}) {
	f := LocalSprintf(key, a)
	PrintGreen(f)
}

func PrintLocalWithYellow(key message.Reference, a ...interface{}) {
	f := LocalSprintf(key, a)
	PrintYellow(f)
}

func PrintLocalWithBlue(key message.Reference, a ...interface{}) {
	f := LocalSprintf(key, a)
	PrintBlue(f)
}

func PrintLocalWithPlain(key message.Reference, a ...interface{}) {
	f := LocalSprintf(key, a)
	PrintPlain(f)
}

func PrintLocalWithRedln(key message.Reference, a ...interface{}) {
	f := LocalSprintf(key, a)
	PrintRedln(f)
}

func PrintLocalWithGreenln(key message.Reference, a ...interface{}) {
	f := LocalSprintf(key, a)
	PrintGreenln(f)
}

func PrintLocalWithYellowln(key message.Reference, a ...interface{}) {
	f := LocalSprintf(key, a)
	PrintYellowln(f)
}

func PrintLocalWithBlueln(key message.Reference, a ...interface{}) {
	f := LocalSprintf(key, a)
	PrintBlueln(f)
}

func PrintLocalWithPlainln(key message.Reference, a ...interface{}) {
	f := LocalSprintf(key, a)
	PrintPlainln(f)
}
