package ircccodes

// IRCommand is essentially a command sent from a remote control to the display
type IRCommand uint

const (
	Button_Display       IRCommand = 5
	Button_Home          IRCommand = 6
	Button_Options       IRCommand = 7
	Button_Return        IRCommand = 8
	Button_Up            IRCommand = 9
	Button_Down          IRCommand = 10
	Button_Right         IRCommand = 11
	Button_Left          IRCommand = 12
	Button_Confirm       IRCommand = 13
	Button_Red           IRCommand = 14
	Button_Green         IRCommand = 15
	Button_Yellow        IRCommand = 16
	Button_Blue          IRCommand = 17
	Button_Num1          IRCommand = 18
	Button_Num2          IRCommand = 19
	Button_Num3          IRCommand = 20
	Button_Num4          IRCommand = 21
	Button_Num5          IRCommand = 22
	Button_Num6          IRCommand = 23
	Button_Num7          IRCommand = 24
	Button_Num8          IRCommand = 25
	Button_Num9          IRCommand = 26
	Button_Num0          IRCommand = 27
	Button_Volume_Up     IRCommand = 30
	Button_Volume_Down   IRCommand = 31
	Button_Mute          IRCommand = 32
	Button_Channel_Up    IRCommand = 33
	Button_Channel_Down  IRCommand = 34
	Button_Subtitle      IRCommand = 35
	Button_DOT           IRCommand = 38
	Button_Picture_Off   IRCommand = 50
	Button_Wide          IRCommand = 61
	Button_Jump          IRCommand = 62
	Button_Sync_Menu     IRCommand = 76
	Button_Forward       IRCommand = 77
	Button_Play          IRCommand = 78
	Button_Rewind        IRCommand = 79
	Button_Prev          IRCommand = 80
	Button_Stop          IRCommand = 81
	Button_Next          IRCommand = 82
	Button_Pause         IRCommand = 84
	Button_Flash_Plus    IRCommand = 86
	Button_Flash_Minus   IRCommand = 87
	Button_TV_Power      IRCommand = 98
	Button_Audio         IRCommand = 99
	Button_Input         IRCommand = 101
	Button_Sleep         IRCommand = 104
	Button_Sleep_Timer   IRCommand = 105
	Button_Video_2       IRCommand = 108
	Button_Picture_Mode  IRCommand = 110
	Button_Demo_Surround IRCommand = 121
	Button_HDMI_1        IRCommand = 124
	Button_HDMI_2        IRCommand = 125
	Button_HDMI_3        IRCommand = 126
	Button_HDMI_4        IRCommand = 127
	Button_Action_Menu   IRCommand = 129
	Button_Help          IRCommand = 130
)
