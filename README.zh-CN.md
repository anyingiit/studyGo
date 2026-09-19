[English](README.md) · **简体中文**

> 英文版是规范版本。本页与 [README.md](README.md) 不一致时，以英文版为准。

<!-- translation-of: README.md sha256:672c73907b06b590 -->

<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# StudyGo

一份个人练习仓库，收录了若干 LeetCode 风格题目的 Go 解法，外加两个从零写的小型数据结构实验；仓库本身没有可安装、可运行的命令行工具。

[![CI](https://github.com/anyingiit/studyGo/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/studyGo/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/studyGo)](LICENSE)

[报告问题](https://github.com/anyingiit/studyGo/issues/new?template=bug_report.yml) · [提出需求](https://github.com/anyingiit/studyGo/issues/new?template=feature_request.yml)

<details>
  <summary>目录</summary>
  <ol>
    <li><a href="#about-the-project">关于本项目</a></li>
    <li><a href="#getting-started">开始使用</a></li>
    <li><a href="#usage">用法</a></li>
    <li><a href="#contributing">参与贡献</a></li>
    <li><a href="#license">许可证</a></li>
    <li><a href="#contact">联系方式</a></li>
  </ol>
</details>

## 关于本项目

studyGo 是 anyingiit 用来练习算法的个人仓库。`function/leetcode/` 下每道 LeetCode 风格题目各占一个解法文件——删除有序数组中的重复项、合并区间、判断括号是否有效、旋转矩阵等等——每个文件都配有一份写在同名文件里的思路说明。`function/study/` 下有两个从零写的小实验：一个基于切片的栈（`myStack.go`），以及一个基于 `sort.Interface` 的选择排序演示（`sortSelice.go`）。`main/main.go` 是一个练习用的入口文件，其函数体几乎全部被注释掉；运行它能编译通过，但不会产生任何输出，仓库里也没有任何部分被真正组装成一个可安装、可直接运行的程序。

计划中的功能与已知问题，见 [open issues](https://github.com/anyingiit/studyGo/issues)。

## 开始使用

### 环境要求

- [Go](https://go.dev/) 1.22 或更高版本——用于构建和测试 `function/` 与 `main/` 下的每个包

### 安装

```sh
git clone https://github.com/anyingiit/studyGo.git
cd studyGo
go build ./...
```

`go build ./...` 不需要拉取任何依赖：该模块只引用标准库，没有第三方依赖需要解析。

## 用法

```sh
go test ./function/study/...
```

这个仓库没有命令行工具可用：`main/main.go` 的 `func main` 函数体几乎全部被注释掉，只剩两行对字符串做切片并丢弃结果的代码，所以 `go run ./main` 编译通过后不会打印任何内容。想看到真实的运行效果，可以运行上面的测试，或者在 `function/leetcode/` 下打开某道题目的 `.go` 文件，旁边就是它的思路说明。

## 参与贡献

欢迎参与。[CONTRIBUTING.md](CONTRIBUTING.md) 说明如何提交 issue 或 pull request，[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 说明对所有参与者的行为要求。

请不要在公开的 issue 或 pull request 中报告安全问题。[SECURITY.md](SECURITY.md) 说明了私下报告的方式。

## 许可证

以 MIT 许可证分发。详见 [LICENSE](LICENSE)。

## 联系方式

项目地址：[https://github.com/anyingiit/studyGo](https://github.com/anyingiit/studyGo)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
