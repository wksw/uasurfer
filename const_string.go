// Code generated by "stringer -type=DeviceType,BrowserName,OSName,Platform -output=const_string.go"; DO NOT EDIT.

package uasurfer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DeviceUnknown-0]
	_ = x[DeviceComputer-1]
	_ = x[DeviceTablet-2]
	_ = x[DevicePhone-3]
	_ = x[DeviceConsole-4]
	_ = x[DeviceWearable-5]
	_ = x[DeviceTV-6]
}

const _DeviceType_name = "DeviceUnknownDeviceComputerDeviceTabletDevicePhoneDeviceConsoleDeviceWearableDeviceTV"

var _DeviceType_index = [...]uint8{0, 13, 27, 39, 50, 63, 77, 85}

func (i DeviceType) String() string {
	if i < 0 || i >= DeviceType(len(_DeviceType_index)-1) {
		return "DeviceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeviceType_name[_DeviceType_index[i]:_DeviceType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BrowserUnknown-0]
	_ = x[BrowserChrome-1]
	_ = x[BrowserIE-2]
	_ = x[BrowserSafari-3]
	_ = x[BrowserFirefox-4]
	_ = x[BrowserAndroid-5]
	_ = x[BrowserOpera-6]
	_ = x[BrowserBlackberry-7]
	_ = x[BrowserUCBrowser-8]
	_ = x[BrowserSilk-9]
	_ = x[BrowserNokia-10]
	_ = x[BrowserNetFront-11]
	_ = x[BrowserQQ-12]
	_ = x[BrowserWechat-13]
	_ = x[BrowserPostman-14]
	_ = x[BrowserFlexPai2-15]
	_ = x[BrowserFlexPai-16]
	_ = x[BrowserMaxthon-17]
	_ = x[BrowserSogouExplorer-18]
	_ = x[BrowserSpotify-19]
	_ = x[BrowserNintendo-20]
	_ = x[BrowserSamsung-21]
	_ = x[BrowserYandex-22]
	_ = x[BrowserCocCoc-23]
	_ = x[BrowserBot-24]
	_ = x[BrowserAppleBot-25]
	_ = x[BrowserBaiduBot-26]
	_ = x[BrowserBingBot-27]
	_ = x[BrowserDuckDuckGoBot-28]
	_ = x[BrowserFacebookBot-29]
	_ = x[BrowserGoogleBot-30]
	_ = x[BrowserLinkedInBot-31]
	_ = x[BrowserMsnBot-32]
	_ = x[BrowserPingdomBot-33]
	_ = x[BrowserTwitterBot-34]
	_ = x[BrowserYandexBot-35]
	_ = x[BrowserCocCocBot-36]
	_ = x[BrowserYahooBot-37]
}

const _BrowserName_name = "BrowserUnknownBrowserChromeBrowserIEBrowserSafariBrowserFirefoxBrowserAndroidBrowserOperaBrowserBlackberryBrowserUCBrowserBrowserSilkBrowserNokiaBrowserNetFrontBrowserQQBrowserWechatBrowserPostmanBrowserFlexPai2BrowserFlexPaiBrowserMaxthonBrowserSogouExplorerBrowserSpotifyBrowserNintendoBrowserSamsungBrowserYandexBrowserCocCocBrowserBotBrowserAppleBotBrowserBaiduBotBrowserBingBotBrowserDuckDuckGoBotBrowserFacebookBotBrowserGoogleBotBrowserLinkedInBotBrowserMsnBotBrowserPingdomBotBrowserTwitterBotBrowserYandexBotBrowserCocCocBotBrowserYahooBot"

var _BrowserName_index = [...]uint16{0, 14, 27, 36, 49, 63, 77, 89, 106, 122, 133, 145, 160, 169, 182, 196, 211, 225, 239, 259, 273, 288, 302, 315, 328, 338, 353, 368, 382, 402, 420, 436, 454, 467, 484, 501, 517, 533, 548}

func (i BrowserName) String() string {
	if i < 0 || i >= BrowserName(len(_BrowserName_index)-1) {
		return "BrowserName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BrowserName_name[_BrowserName_index[i]:_BrowserName_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OSUnknown-0]
	_ = x[OSWindowsPhone-1]
	_ = x[OSWindows-2]
	_ = x[OSMacOSX-3]
	_ = x[OSiOS-4]
	_ = x[OSAndroid-5]
	_ = x[OSBlackberry-6]
	_ = x[OSChromeOS-7]
	_ = x[OSKindle-8]
	_ = x[OSWebOS-9]
	_ = x[OSLinux-10]
	_ = x[OSPlaystation-11]
	_ = x[OSXbox-12]
	_ = x[OSNintendo-13]
	_ = x[OSBot-14]
}

const _OSName_name = "OSUnknownOSWindowsPhoneOSWindowsOSMacOSXOSiOSOSAndroidOSBlackberryOSChromeOSOSKindleOSWebOSOSLinuxOSPlaystationOSXboxOSNintendoOSBot"

var _OSName_index = [...]uint8{0, 9, 23, 32, 40, 45, 54, 66, 76, 84, 91, 98, 111, 117, 127, 132}

func (i OSName) String() string {
	if i < 0 || i >= OSName(len(_OSName_index)-1) {
		return "OSName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OSName_name[_OSName_index[i]:_OSName_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PlatformUnknown-0]
	_ = x[PlatformWindows-1]
	_ = x[PlatformMac-2]
	_ = x[PlatformLinux-3]
	_ = x[PlatformiPad-4]
	_ = x[PlatformiPhone-5]
	_ = x[PlatformiPod-6]
	_ = x[PlatformBlackberry-7]
	_ = x[PlatformWindowsPhone-8]
	_ = x[PlatformPlaystation-9]
	_ = x[PlatformXbox-10]
	_ = x[PlatformNintendo-11]
	_ = x[PlatformBot-12]
}

const _Platform_name = "PlatformUnknownPlatformWindowsPlatformMacPlatformLinuxPlatformiPadPlatformiPhonePlatformiPodPlatformBlackberryPlatformWindowsPhonePlatformPlaystationPlatformXboxPlatformNintendoPlatformBot"

var _Platform_index = [...]uint8{0, 15, 30, 41, 54, 66, 80, 92, 110, 130, 149, 161, 177, 188}

func (i Platform) String() string {
	if i < 0 || i >= Platform(len(_Platform_index)-1) {
		return "Platform(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Platform_name[_Platform_index[i]:_Platform_index[i+1]]
}
