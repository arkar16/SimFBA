package util

import (
	"math"
	"strconv"
	"strings"
)

func GetPassStatement(yards int, offensiveFormation, playName, poa, recLabel string,
	touchdown, outOfBounds, twoPtConversion, fumble,
	safety, scramble, sack, complete, interception bool,
	turnOverLabel string) string {
	snapText := getSnapText(offensiveFormation)
	scrambleText := ""
	if scramble {
		scrambleText = getScrambleText(yards, touchdown)
		return snapText + scrambleText
	}
	if sack {
		sackText := getSackText(safety, fumble, touchdown)
		return snapText + sackText
	}
	if interception {
		intText := getInterceptText(yards, recLabel, turnOverLabel, fumble, touchdown)
		return intText
	}
	if complete {

	} else {

	}

	finalString := snapText + scrambleText
	return finalString
}

func getInterceptText(yards int, recLabel, turnOverLabel string, fumble, touchdown bool) string {
	absYards := math.Abs(float64(yards))
	yardsInt := int(absYards)
	ydStr := strconv.Itoa(yardsInt)
	yardsStr := GetYardsString(int8(yardsInt))
	var list []string
	// Very rare case -- the problem is that I don't think we have the capacity to tell who scored based on the play data tangible
	if fumble && touchdown {
		list = append(list, " and he's picked off! Caught by "+turnOverLabel+" with the catch and he makes a run for it! Brought down, an- the ball is lose! It looks like it's a fight for it, and it's picked up! The player's making it to the endzone! TOUCHDOWN! ")
	} else if fumble && !touchdown {
		list = append(list, " and he's picked off! Caught by "+turnOverLabel+" with the catch and he makes a run for it! Brought down, an- the ball is lose! It's a fight for the pigskin! ")
	} else if !fumble && touchdown {
		list = append(list, " and he's picked off! Caught by "+turnOverLabel+" with the catch and he makes a return all the way into the endzone! TOUCHDOWN! ")
	} else {
		list = append(list, " and he's picked off! Caught by "+turnOverLabel+" with the catch and he makes a return for "+ydStr+yardsStr)
		if yards > 19 {

		} else if yards > 9 {

		} else if yards > 4 {

		} else {

		}
	}

	return PickFromStringList(list)
}

func getSackText(safety, fumble, touchdown bool) string {
	var list []string
	if safety {
		list = append(list,
			"Is sacked in the end zone, resulting in a safety! ",
			"Can't escape the grasp of the defenders, leading to a safety with that sack! ",
			"Is taken down in the end zone for a safety, a crucial play! ",
			"The pressure gets to him, resulting in a sack for a safety! ",
			"Is overwhelmed by the defense and sacked for a safety! ")
	} else if fumble && !touchdown {
		list = append(list,
			"Tries to evade the rush but is sacked! What's this? The ball is loose! It's a scramble to recover the football! ",
			"Takes too long to find a man and the defense has broken through! A sack on the play -- and the ball has fumbled! Both sides are scrambling to recover the ball. ",
			"Tries to throw it away but is sacked on the play! And wi- the ball is loose onto the field! Everyone's scrambling for the ball in an attempt to recover it. ",
			"Can't find a man and is sacked AND fumbles the ball! The defense is scrambling to recover it! ",
			"Can't evade the rush and is sack. What's this? He's lost his grip while being sacked, and the ball is loose on the field! ",
			"And he's brought down by the pass rush. The hit caused a fumble! The quarterback loses the ball as he's sacked! ",
			"The pocket collapses, and he's sacked! What's this? The ball is knocked loose! A potential turnover here! ",
			"The defense has broken through and he's taken down hard - and the ball pops out! A fumble during the sack! Both sides are trying to recover! ",
		)
	} else if fumble && touchdown {
		list = append(list,
			"Can't find a man and is sacked on the play! What's this? There's a fumble on the field! The defense has scooped up the ball and is going... all the way! TOUCHDOWN! ",
			"The defense has broken through has sacked the quarterbac- and the ball is fumbled! It's been scooped up by the defense and is being returned. Not a man in sight -- that's a TOUCHDOWN! ",
			"Hesitates on a throw and is sacked. An- what's this? The ball is loose! It's a disaster for the offense as the defense has scooped up the ball and it's been taken back for a TOUCHDOWN! ",
			"The pocket collapses, and he's sacked! What's this? The ball has fumbled on the field! The defense capitalizes on the fumble with the recovery and a return to the endzone - TOUCHDOWN! ",
			"Takes too long on the throw and is sacked - the ball comes loose! The defense has scooped it up and is making a return for the endzone! TOUCHDOWN! ",
		)
	} else {
		list = append(list, "Hesitates on throwing the ball and is sacked on the play! ",
			"Tries to look for an open man, but is brought down by the defense. A huge sack on the play! ",
			"Can't find an open receiver and is sacked behind the line! ",
			"The pocket collapses, and he's sacked! ",
			"The pocket collapses, and he can't scramble out. A sack on the play! ",
			"What's this? The rush has overwhelmed the offensive line and the QB's taken down for a sack! ",
			"Tries to evade the rush but is sacked! ",
			"Takes too long to find a man and the defense has broken through! A sack on the play! ",
			"Tries to throw it away but is sacked on the play! ",
			"Is wrapped up and sacked, a significant loss on the play! ",
			"Faces a fierce pass rush and is sacked, thwarting the drive! ",
			"The defense breaks through and he's sacked, a big play! ",
			"Attempts to scramble but is caught and sacked! ",
			"Holds onto the ball too long and is sacked by the oncoming defenders! ",
			"Is hit and sacked, the defensive line breaking through! ",
		)
	}
	return PickFromStringList(list)
}

func getScrambleText(yards int, touchdown bool) string {
	gainStatement := getGainSuffix(yards > 0, yards)
	yardsStr := GetYardsString(int8(yards))
	if !touchdown {
		list := []string{"Leaves the pocket on and scrambles for " + gainStatement,
			"Sees pressure and scrambles out of the pocket for" + gainStatement + yardsStr,
			"Scrambles out of the pocket for" + gainStatement + yardsStr,
			"Leaves the pocket on a scramble, evading defenders for" + gainStatement + yardsStr,
			"Tucks the ball and scrambles, looking for open space. Runs for" + gainStatement + yardsStr,
			"Dodges the rush and scrambles out to the side for" + gainStatement + yardsStr,
			"Takes evasive action and scrambles to avoid a sack," + gainStatement + yardsStr,
			"Finds no one open and decides to scramble for yardage," + gainStatement + yardsStr,
			"Breaks away from pressure, scrambles for" + gainStatement + yardsStr,
			"Under duress, elects to scramble out of the pocket for" + gainStatement + yardsStr,
			"Sees the pocket collapsing and takes off on a scramble. Running for" + gainStatement + yardsStr,
			"Sees an opening and quickly scrambles out of the pocket for" + gainStatement + yardsStr,
			"Avoids the sack with a quick scramble to the side," + gainStatement + yardsStr,
			"Uses his legs to escape the pocket, runs for" + gainStatement + yardsStr,
			"Finds a lane and scrambles to exploit the gap, runs for" + gainStatement + yardsStr,
		}
		return PickFromStringList(list)
	}
	list := []string{"Leaves the pocket on and scrambles into the endzone for the TOUCHDOWN! ",
		"Sees pressure and scrambles out of the pocket gets right into the endzone! TOUCHDOWN! ",
		"Scrambles out of the pocket and makes a dive right into the endzone! TOUCHDOWN! ",
		"Leaves the pocket on a scramble, evading defenders on the goalline and makes it into the endzone - TOUCHDOWN! ",
		"Tucks the ball and scrambles for an open gap on the goalline. Dives into the endzone for the TOUCHDOWN! ",
		"Takes evasive action and scrambles to avoid a sack, runs to the edge of the goalline and makes it in - TOUCHDOWN! ",
		"Finds no one open and decides to scramble for the endzone. Succeeds, it's a TOUCHDOWN! ",
		"Sees the pocket collapsing and takes off on a scramble. Runs for the goalline and makes it into the endzone - TOUCHDOWN! ",
		"Avoids the sack with a quick scramble right into the endzone - TOUCHDOWN! ",
	}
	return PickFromStringList(list)
}

func getSnapText(form string) string {
	list := []string{" takes the snap. "}
	isShotgunPlay := CheckSubstring(form, "Gun")
	if isShotgunPlay {
		list = append(list, " drops back to pass. ",
			" takes the snap in the shotgun formation. ",
			" gets the ball in the shotgun.",
			" fields the snap in the shotgun.",
			" catches the snap while in the shotgun, scanning for options.",
			" from the shotgun, secures the snap and setups the throw.",
		)
	} else {
		list = append(list, " takes the snap from under center. ",
			" takes the ball from under center. ",
			" gets the snap and drops back. ",
			" receives the snap under center. Looks to pass. ",
			" under center, takes the snap and looks to pass. ",
			" collects the snap under center and prepares to throw. ",
			" snaps up the ball from under center and eyes his targets. ")
	}

	return PickFromStringList(list)
}

func getThrowingVerb(yards int) string {
	list := []string{"throws", "slings", "makes the pass", "fires", "lobs", "hurls",
		"tosses", "flings"}
	if yards > 19 {
		list = append(list, "chucks")
	}
	return PickFromStringList(list)
}

func getDistance(playName string, yards int) string {
	list := []string{}
	direction := GenerateIntFromRange(1, 3) // 1 == left, 2 == Middle, 3 == right
	dirs := ""
	dirsList := []string{}
	if direction == 1 {
		dirsList = []string{"left side", "left"}
		dirs = "left side"
	} else if direction == 2 {
		dirsList = []string{"middle of the field", "middle"}
		dirs = "middle of the field"
	} else {
		dirsList = []string{"right side", "right"}
	}
	dirs = PickFromStringList(dirsList)
	if yards > 45 {
		list = append(list, "across midfield")
		if direction == 1 || direction == 3 {
			list = append(list, "across the field to the "+dirs, "across midfield to the "+dirs)
		}
	}
	if yards > 19 {

	}

	return PickFromStringList(list)
}

func CheckSubstring(text, subtext string) bool {
	return strings.Contains(text, subtext)
}

func GetYardsString(yds int8) string {
	yards := " yards. "
	if yds == 1 || yds == -1 {
		yards = " yard. "
	}
	return yards
}
