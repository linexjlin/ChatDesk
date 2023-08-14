// 导入需要的头文件
#import <Cocoa/Cocoa.h>
#import <WebKit/WebKit.h>
#include "webview.h"

// 声明全局变量
NSWindowController *windowController = nil;
NSWindow *window = nil;
WKWebView *webView = nil;

// 配置应用窗口
void configureAppWindow(char* title, int width, int height)
{
  // 如果窗口控制器已经存在，那么就不再进行配置。
  if (windowController != nil) {
    return;
  }

  // 获取应用实例
  NSApplication *app = [NSApplication sharedApplication];
  // 将应用的激活策略设置为辅助应用，以避免应用出现在 Dock 中
  [app setActivationPolicy:NSApplicationActivationPolicyAccessory]; 
  [app activateIgnoringOtherApps:YES];

  // 创建窗口大小和位置的矩形
  NSRect frame = NSMakeRect(0, 0, width, height);
  // 定义窗口的样式，包括标题栏、可调整大小、可关闭和可最小化
  int mask = NSWindowStyleMaskTitled | NSWindowStyleMaskResizable | NSWindowStyleMaskClosable | NSWindowStyleMaskMiniaturizable;
  // 初始化窗口
  window = [[NSWindow alloc] initWithContentRect:frame
                              styleMask:mask
                              backing:NSBackingStoreBuffered
                              defer:NO];
  // 设置窗口标题
  [window setTitle:[[NSString alloc] initWithUTF8String:title]];
  // 将窗口置于屏幕中央
  [window center];

  // 获取窗口的内容视图，并在其中添加一个 WebView
  NSView *contentView = [window contentView];
  webView = [[WKWebView alloc] initWithFrame:[contentView bounds]];
  // 添加 WebView 到内容视图中，并设置约束使其填满整个内容视图
  [webView setTranslatesAutoresizingMaskIntoConstraints:NO];
  [contentView addSubview:webView];
  [contentView addConstraint:
    [NSLayoutConstraint constraintWithItem:webView
        attribute:NSLayoutAttributeWidth
        relatedBy:NSLayoutRelationEqual
        toItem:contentView
        attribute:NSLayoutAttributeWidth
        multiplier:1
        constant:0]];
  [contentView addConstraint:
    [NSLayoutConstraint constraintWithItem:webView
        attribute:NSLayoutAttributeHeight
        relatedBy:NSLayoutRelationEqual
        toItem:contentView
        attribute:NSLayoutAttributeHeight
        multiplier:1
        constant:0]];

  // 创建窗口控制器，并与窗口关联
  windowController = [[NSWindowController alloc] initWithWindow:window];
  
  // 释放标题字符串占用的内存
  free(title);
  // 开始应用的事件循环
  [NSApp run];
}

// 在主线程中显示应用窗口，并加载指定 URL 的网页
void doShowAppWindow(char* url)
{
  // 如果窗口控制器不存在，那么就无法打开窗口
  if (windowController == nil) {
    return;
  }

  // 创建 URL 和请求对象
  id nsURL = [NSURL URLWithString:[[NSString alloc] initWithUTF8String:url]];
  id req = [[NSURLRequest alloc] initWithURL: nsURL
                                 cachePolicy: NSURLRequestUseProtocolCachePolicy
                                 timeoutInterval: 5];
  // 在 WebView 中加载请求
  [webView loadRequest:req];
  // 显示窗口
  [windowController showWindow:window];
  // 释放 URL 字符串占用的内存
  free(url);
}

// 在主线程中调用 doShowAppWindow 方法
void showAppWindow(char* url)
{
  dispatch_async(dispatch_get_main_queue(), ^{
    doShowAppWindow(url```objectc
  });
}

// 最小化应用窗口
void minimizeAppWindow()
{
  // 如果窗口不存在，那么就无法最小化窗口
  if (window == nil) {
    return;
  }

  // 在主线程中最小化窗口
  dispatch_async(dispatch_get_main_queue(), ^{
    [window miniaturize:nil];
  });
}