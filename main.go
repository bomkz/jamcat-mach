package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	fmt.Println(
		`This program should be able to find where VTOL VR is located, even if you installed steam or VTOL VR in a drive other than C:\.
However, there might be some edge cases where this wouldn't work. Just message me on Discord, or file an issue on Github so I can see if I can fix it for you.

GLWT(Good Luck With That) Public License
Copyright (c) Everyone, except Author

Everyone is permitted to copy, distribute, modify, merge, sell, publish,
sublicense or whatever they want with this software but at their OWN RISK.

		   Preamble

The author has absolutely no clue what the code in this project does.
It might just work or not, there is no third option.


GOOD LUCK WITH THAT PUBLIC LICENSE
TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION, AND MODIFICATION

0. You just DO WHATEVER YOU WANT TO as long as you NEVER LEAVE A
TRACE TO TRACK THE AUTHOR of the original product to blame for or hold
responsible.

IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
DEALINGS IN THE SOFTWARE.

Good luck and Godspeed.

------------------------------------------------------------------------------------------------------
		Jet Audio and Music Control Access Terminal-Media Access and Control Hub release v1.0
		
JAMCAT-MACH-one made by @f45a - discord.
GIthub repo: https://github.com/angelfluffyookami/jet-mach
Please be aware, this utility will append ".bkp" to any mp3 files currently in the VTOL VR folder.
If closed properly without crashing or using TaskMan, will remove ".bkp" from said files and delete 
any temporary files placed in this folder.

I have not experienced any data loss using this, and there shouldn't be any to expect, however, please don't be stupid:
		- Don't touch the Player.log file while program is in use.
		- Do NOT modify, add, or remove any file in VTOLVR's RadioMusic folder while in use. 
			(I did add a way to handle this, however, I cannot guarantee it will work 100 percent of the time, so just d o n t)`)

	bkpPlayerMp3()
	// InitMP3()

	go readLog()
	fmt.Println("VTOL VR Windows Media Control Bridge is nor listening to log events.")

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"revertBkp": func(ctx context.Context) error {
			revertBkp()
			return nil
		},
	})

	<-wait
}
