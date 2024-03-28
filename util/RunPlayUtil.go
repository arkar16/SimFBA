package util

func GetRunVerb(yards int, playName string, touchdown, outOfBounds, twoPtConversion, fumble, safety bool) string {
	list := []string{}
	if !touchdown && !outOfBounds && !twoPtConversion && !fumble && !safety {
		list = append(list, " carries for ", " runs the ball ", " runs it for ", " runs it by for ", " advances the ball for ", " hustles for ")
	}
	// Basic run
	if playName == "Inside Run Left" || playName == "Inside Run Right" {
		list = append(list, getInsideRunList(yards, playName == "Inside Run Left", touchdown, outOfBounds, twoPtConversion, fumble, safety)...)
	} else if playName == "Outside Run Left" || playName == "Outside Run Right" {
		list = append(list, getOutsideRunList(yards, playName == "Outside Run Left", touchdown, outOfBounds, twoPtConversion, fumble, safety)...)
	} else if playName == "Power Run Left" || playName == "Power Run Right" {
		list = append(list, getPowerRunList(yards, playName == "Power Run Left", touchdown, outOfBounds, twoPtConversion, fumble, safety)...)
	} else if playName == "Draw Left" || playName == "Draw Right" {

	} else if playName == "Read Option Left" || playName == "Read Option Right" {

	} else if playName == "Speed Option Left" || playName == "Speed Option Right" {

	} else if playName == "Inverted Option Left" || playName == "Inverted Option Right" {

	} else if playName == "Choice Inside Left" || playName == "Choice Inside Right" {

	} else if playName == "Choice Outside Left" || playName == "Choice Outside Right" {

	} else if playName == "Peek Inside Left" || playName == "Peek Inside Right" {

	} else if playName == "Peek Outside Left" || playName == "Peek Outside Right" {

	} else if playName == "Peek Power Left" || playName == "Peek Power Right" {

	}

	return PickFromStringList(list)
}

func getPowerRunList(yards int, isleft, touchdown, outOfBounds, twoPtConversion, fumble, safety bool) []string {
	list := []string{}
	direction := "left"
	if isleft {
		direction = "right"
	}
	// Depending on the yards gained or lost, append appropriate descriptions
	if touchdown {
		list = append(list, " bulldozes through the line and powers his way into the end zone for a touchdown! ",
			" charges forward, pushing defenders back and breaking the plane for a score! ",
			" showcases sheer strength, fighting off tacklers and plunging into the end zone! ",
			" demonstrates a powerful run, shedding tackles and reaching paydirt! ",
			" steamrolls his way through the defense and into the end zone for six! ",
			" plows over the goal line, dragging defenders with him for the touchdown! ")
	} else if outOfBounds {
		list = append(list, " makes a dash to the "+direction+" and steps out after gaining ",
			" stretches the play to the "+direction+" before being forced out, gaining ",
			" barrels to the "+direction+" and is pushed out after a solid gain of ",
			" powers his way to the sideline, gaining valuable yards before stepping out ",
			" shows off his strength, pushing through to the "+direction+" before going out at ",
			" maintains a strong run to the "+direction+" and exits the field after gaining ")
	} else if twoPtConversion && touchdown {
		list = append(list, " pushes through the line for the crucial two-point conversion! ",
			" barrels into the end zone for the two-point conversion! ",
			" dives across the goal line, and the two-point attempt is good! ",
			" finds a tiny crease and converts for two points! ",
			" outmuscles the defenders and powers in for the conversion! ",
			" drives through the defenders and secures the two-point conversion with power! ",
			" leans into the defense, pushing forward for the successful conversion! ",
			" fights through the line with a surge of strength for two points! ",
			" delivers a forceful run, crossing into the end zone for the conversion! ")
	} else if twoPtConversion && !touchdown {
		list = append(list, " can't push through the line for the crucial two-point conversion! No good! ",
			" barrels into a lineman and misses the end zone for the two-point conversion! No good! ",
			" makes a dive, but misses the goal line. The two-point attempt is no good! ",
			" can't find a gap and is brought down. Two point conversion is no good! ",
			" is outmuscled by the defenders. The two point conversion is no good! ",
			" attempts to bulldoze through but is halted before the line. No conversion! ",
			" tries to overpower the defense but falls short of the end zone! ",
			" pushes with all his might but can't break the plane. Conversion fails! ",
			" meets a wall of defenders and can't punch it through for the two points! ")
	} else if fumble {
		list = append(list, " fights for extra yards but loses the ball in the skirmish after gaining ",
			" powers forward but loses control of the football, fumbling after ",
			" has the ball jarred loose after a hard-earned ",
			" makes a costly mistake, fumbling the ball away after ",
			" can't hold on to the pigskin, and it's on the ground after ")
	} else if safety {
		list = append(list, " is caught in the end zone for a safety! ",
			" is wrapped up in the end zone, and it's two points the other way! ",
			" gets trapped in the end zone during a power run, resulting in a safety! ",
			" can't break free from the defense in the end zone, conceding a safety! ",
			" is swarmed by the defense for a safety! ",
			" tries to get out of the end zone but is trapped for a safety! ",
			" can't escape the clutches of the defense – that's a safety! ")
	} else {
		switch {
		case yards > 14:
			list = append(list, " bursts through the "+direction+" and breaks free for a massive gain of ",
				" passes the line of scrimmage, dodges right past a linebacker and makes a break for it before being taken down for ",
				" charges upfield, shrugging off tacklers and marking a significant run of ",
				" blazes down the field, eluding defenders and finally tackled after a fantastic run of ",
				" blazes down the field, eluding defenders and finally tackled. Fantastic gain of ",
				" chugs downfield like a freight train, eventually brought down after a tremendous run of ",
				" gallops downfield like a gazelle, eventually brought down after a tremendous effort. A run for ",
				" steamrolls through defenders, showcasing brute strength for a gain of ",
				" bulldozes his way downfield, an unstoppable force for ",
				" plows through the line like a hot knife through butter, amassing ",
				" demonstrates sheer power, breaking multiple tackles on his way to a gain of ",
				" turns a routine power run into a highlight reel, thundering down the field for ",
				" embodies the power back archetype, trampling over defenders for ")
		case yards > 9:
			list = append(list,
				" powers through the "+direction+" for a strong run, picking up ",
				" finds a hole in the "+direction+" and hits it with authority, chalking up ",
				" darts through the "+direction+" side, with a burst of speed, racking up ",
				" carves up the defense on the "+direction+" for a healthy gain of ",
				" exploits a gap to the "+direction+" and rumbles past defenders for ",
				" turns on the afterburners and leaves the defense in the dust, a solid run for ",
				" weaves through the defense to the "+direction+" and breaks away, tallying ",
				" finds daylight on the "+direction+", turning on the jets for a gain of ",
				" blasts past the line into the "+direction+" side, barreling forward for ",
				" rumbles ahead with determination, pushing defenders aside for ",
				" showcases a blend of power and agility, turning a tough run into ",
				" powers ahead, dragging defenders with him for a solid gain of ",
				" barrels down the "+direction+", shrugging off tackles for ",
				" exhibits a classic power run, plowing forward to pick up ",
				" charges ahead with sheer force, carving out a nice chunk of yards for ")
		case yards > 4:
			list = append(list,
				" muscles his way past the line and into the "+direction+", securing ",
				" surges forward with power, shouldering through to the "+direction+" for ",
				" follows his blockers to the "+direction+" and powers through for ",
				" with a hop and a skip, sidesteps a defender and sprints to the "+direction+" for ",
				" hits the hole on the "+direction+" with gusto, picking up ",
				" grinds out a hard-earned few yards, battling his way through for ",
				" fights for every inch, muscling ahead for ",
				" pushes the pile, demonstrating strength as he gains ",
				" gains some tough yards on the power run for ",
				" moves the chains bit by bit, a testament to his strength, netting ",
				" battles through the line, securing a few valuable yards for ")
		case yards > 0:
			list = append(list,
				" nudges forward on the "+direction+", managing to get ",
				" makes a run for it but is brought down quickly. Ran for ",
				" speeds past a defender but is swiftly tackled. Ran for ",
				" scoots up the field for ",
				" squeezes past the line of scrimmage but is quickly halted. A gain of ",
				" charges forward but meets a wall of defenders. A short gain of ",
				" meets a wall of defenders, eking out just ",
				" barely nudges forward, a testament to the defense's strength, gaining ",
				" finds little room and is stopped after a modest gain of ",
				" stumbles slightly but regains footing, a short return for ",
				" gets a grip on the ball and scurries ahead, only managing ",
				" gathers the kick and charges ahead into the arms of the opposition, a minimal gain of ")
		case yards == 0:
			list = append(list, " is met at the line of scrimmage, no gain on the play. ",
				" is immediately stopped at the line of scrimmage, no gain on the play. ",
				" is stonewalled at the line, no gain on the play. ",
				" finds no room to maneuver, stopped dead in his tracks. ",
				" is immediately tackled by the defensive line. No gain on the play. ")
		default: // Negative yardage
			list = append(list, " is wrapped up behind the line, a loss of ",
				" is stopped in the backfield for ",
				" meets a defender and is tackled in the backfield for ",
				" tries to find a gap, and just like that he's stopped for ",
				" is swallowed by the defense, losing ",
				" can't find a crease and is taken down behind the line for a loss of ",
				" attempts to find a gap and is brought down for ")
		}
	}
	return list
}

func getOutsideRunList(yards int, isleft, touchdown, outOfBounds, twoPtConversion, fumble, safety bool) []string {
	list := []string{}
	direction := "left"
	if isleft {
		direction = "right"
	}
	// Depending on the yards gained or lost, append appropriate descriptions
	if touchdown {
		list = append(list, " runs outside the line and bursts into the end zone for a touchdown! ",
			" dashes around the defense and into the promised land for six! Ran ",
			" avoids the traffic and dances into the end zone! Ran ",
			" counters around line and sprints into the end zone untouched! Ran ",
			" turns on the speed and leaves everyone in his dust – that's a touchdown! Ran ",
			" breaks an outside tackle, then another, and he's in for the score! Ran ")
	} else if outOfBounds {
		list = append(list, " makes a dash to the "+direction+" and steps out after gaining ",
			" skirts the sideline before stepping out, adding ",
			" scampers to the edge and just tiptoes out of bounds, picking up ",
			" races towards the sideline, securing a few extra yards before going out at ",
			" stretches the play to the "+direction+" before being forced out, gaining ",
			" finds a lane to the "+direction+" but is pushed out after a solid gain of ",
			" accelerates to the "+direction+" and is nudged out, notching up ",
			" cuts to the "+direction+", evading a tackle and stepping out with a gain of ",
			" pushes his way to the "+direction+" sideline and steps out, bagging ",
			" showcases agility with a dash to the "+direction+", stepping out after collecting ")
	} else if twoPtConversion && touchdown {
		list = append(list, " pushes around the line for the crucial two-point conversion! ",
			" barrels into the end zone for the two-point conversion! ",
			" makes a run for the goal line, and the two-point attempt is good! ",
			" avoids the dogpile and converts for two points! ",
			" outruns the defenders on the outside and powers in for the conversion! ",
			" finds a crease to the outside and crosses into the end zone for the two-pointer! ",
			" sweeps to the "+direction+" and breaks the plane for a crucial two points! ",
			" darts to the "+direction+" edge and secures the two-point conversion with ease! ",
			" capitalizes on the outside track and dives in for the additional two points! ")
	} else if twoPtConversion && !touchdown {
		list = append(list, " can't find a gap on the outside for the crucial two-point conversion! No good! ",
			" barrels into a lineman and misses the end zone for the two-point conversion! No good! ",
			" makes an outside dive, but misses the goal line. The two-point attempt is no good! ",
			" is stopped outside the line and is brought down. Two point conversion is no good! ",
			" is bottlenecked on the outside route, failing to convert the two points! ",
			" tries an outside leap but falls short of the end zone, two points denied! ",
			" is cornered to the "+direction+" and stopped, missing the critical two points! ",
			" aims for the corner but is tackled just shy of the goal line, no conversion! ")
	} else if fumble {
		list = append(list, " coughs up the ball after a gain of ",
			" loses the grip on the ball mid-run after advancing ",
			" has the ball jarred loose after a hard-earned ",
			" makes a costly mistake, fumbling the ball away after ",
			" can't hold on to the pigskin, and it's on the ground after ",
			" navigates to the "+direction+" but loses control of the ball, dropping it after ",
			" is on a promising run to the "+direction+" but fumbles, the ball slips away after ",
			" makes a break to the "+direction+" but the ball is knocked loose, fumbling after ")
	} else if safety {
		list = append(list, " is caught in the end zone for a safety! ",
			" is wrapped up in the end zone, and it's two points the other way! ",
			" is swarmed by the defense for a safety! ",
			" tries to get out of the end zone but is trapped for a safety! ",
			" can't escape the clutches of the defense – that's a safety! ",
			" is ensnared in the end zone while trying an outside run, resulting in a safety! ",
			" attempts to escape the end zone via the "+direction+" but is tackled for a safety! ",
			" gets cornered in the end zone on an outside attempt, conceding two points! ",
			" retreats to the "+direction+" in the end zone but is caught, a costly safety! ")
	} else {
		switch {
		case yards > 14:
			list = append(list, " bursts through the "+direction+" and breaks free for a massive gain of ",
				" spins past a defender, runs toward the "+direction+" side and makes some headway! Bulldozes a defender but is quickly pulled down. Wow. Ran for ",
				" spins past a defender, runs toward the "+direction+" side. Bulldozes a defender but is quickly pulled down. Wow. Ran for ",
				" charges upfield, shrugging off tacklers and marking a significant run of ",
				" blazes down the field, eluding defenders and finally tackled after a fantastic run of ",
				" blazes down the field, eluding defenders and finally tackled. Fantastic gain of ",
				" gallops downfield like a freight train, eventually brought down after a tremendous run of ",
				" turns on the gas and speeds off, eventually brought down after a tremendous effort. A run for ")
		case yards > 9:
			list = append(list,
				" darts through the "+direction+" side, with a burst of speed, racking up ",
				" carves up the defense on the outside "+direction+" for a healthy gain of ",
				" runs outside to the "+direction+" and rumbles past defenders for ",
				" turns on the afterburners and leaves the defense in the dust, a solid run for ",
				" weaves through the defense to the "+direction+" and breaks away, tallying ",
				" finds daylight on the "+direction+", turning on the jets for a gain of ",
				" runs around the line into the "+direction+" side, barreling forward for ")
		case yards > 4:
			list = append(list,
				" muscles his way around the line and into the "+direction+", securing ",
				" takes a handoff to the outside "+direction+" and barrels ahead for ",
				" follows his blockers to the outside "+direction+" and speeds through for ",
				" counters to the "+direction+" with gusto, picking up ",
				" accelerates around the edge, turning upfield for a gain of ",
				" rounds the "+direction+" corner with a burst, picking up ",
				" glides to the "+direction+", using his speed to notch ",
				" skirts the "+direction+" sideline, eluding tackles for ",
				" sweeps to the "+direction+", cutting upfield for ",
				" darts around the "+direction+" side, scampering for ",
				" showcases his agility, turning the corner on the "+direction+" for ",
			)
		case yards > 0:
			list = append(list,
				" nudges forward on the "+direction+", managing to get ",
				" edges his way to the "+direction+" for a short gain of ",
				" maneuvers around the "+direction+" end, eking out ",
				" pushes around the "+direction+", scraping together ",
				" stretches the play to the "+direction+", barely snagging ",
				" battles past the "+direction+" edge for a modest ",
				" skirts the fringe, squeezing out a few yards to the "+direction+" for ",
				" leans forward, fighting for every inch to the "+direction+", securing ",
				" makes a run for it but is brought down quickly. Ran for ",
				" speeds past a defender but is swiftly tackled. Ran for ",
				" finds little room on the outside and is stopped after a modest gain of ",
				" stumbles slightly but regains footing, a short return for ",
				" gets a grip on the ball and scurries ahead, only managing ",
			)
		case yards == 0:
			list = append(list, " is met on the outside, no gain on the play. ",
				" is immediately stopped on the outside, no gain on the play. ",
				" can't edge past the defensive line. No gain on the play. ",
				" tries to turn the corner on the "+direction+" but is stopped dead in his tracks, no gain.",
				" is corralled at the "+direction+" boundary, unable to advance, no gain.",
				" attempts an outside run to the "+direction+" but is swarmed, no progress made. ",
				" seeks an opening on the "+direction+" but is stymied at the line, no gain. ")
		default: // Negative yardage
			list = append(list, " is wrapped up behind the line, a loss of ",
				" is stopped in the backfield for ",
				" meets a defender and is tackled in the backfield for ",
				" is stopped outside the line and brought backwards for ",
				" attempts to find a gap and is brought down for ",
				" is cornered and taken down behind the line, losing ",
				" attempts to swing to the "+direction+" but is dropped for a loss of ",
				" is ensnared by the defense on the "+direction+", a setback of ",
				" tries to make headway to the "+direction+" but retreats, losing ",
				" is overwhelmed in the backfield trying to go "+direction+", a loss of ",
				" can't turn upfield on the "+direction+" and is tackled for a loss of ")
		}
	}
	return list
}

func getInsideRunList(yards int, isleft, touchdown, outOfBounds, twoPtConversion, fumble, safety bool) []string {
	list := []string{}
	direction := "left"
	if isleft {
		direction = "right"
	}
	// Depending on the yards gained or lost, append appropriate descriptions
	if touchdown {
		list = append(list, " finds a seam and bursts into the end zone for a touchdown! ",
			" slices through the defense and into the promised land for six! Ran ",
			" zigzags his way through the traffic and dances into the end zone! Ran ",
			" blasts through the line and sprints into the end zone untouched! Ran ",
			" turns on the speed and leaves everyone in his dust – that's a touchdown! Ran ",
			" breaks one tackle, then another, and he's in for the score! Ran ")
	} else if outOfBounds {
		list = append(list, " makes a dash to the "+direction+" and steps out after gaining ",
			" skirts the sideline before stepping out, adding ",
			" scampers to the edge and just tiptoes out of bounds, picking up ",
			" races towards the sideline, securing a few extra yards before going out at ",
			" stretches the play to the "+direction+" before being forced out, gaining ")
	} else if twoPtConversion && touchdown {
		list = append(list, " pushes through the line for the crucial two-point conversion! ", " barrels into the end zone for the two-point conversion! ",
			" dives across the goal line, and the two-point attempt is good! ",
			" finds a tiny crease and converts for two points! ",
			" outmuscles the defenders and powers in for the conversion! ")
	} else if twoPtConversion && !touchdown {
		list = append(list, " can't push through the line for the crucial two-point conversion! No good! ",
			" barrels into a lineman and misses the end zone for the two-point conversion! No good! ",
			" makes a dive, but misses the goal line. The two-point attempt is no good! ",
			" can't find a gap and is brought down. Two point conversion is no good! ",
			" is outmuscled by the defenders. The two point conversion is no good! ")
	} else if fumble {
		list = append(list, " coughs up the ball after a gain of ",
			" loses the grip on the ball mid-run after advancing ",
			" has the ball jarred loose after a hard-earned ",
			" makes a costly mistake, fumbling the ball away after ",
			" can't hold on to the pigskin, and it's on the ground after ")
	} else if safety {
		list = append(list, " is caught in the end zone for a safety! ",
			" is wrapped up in the end zone, and it's two points the other way! ",
			" is swarmed by the defense for a safety! ",
			" tries to get out of the end zone but is trapped for a safety! ",
			" can't escape the clutches of the defense – that's a safety! ")
	} else {
		switch {
		case yards > 14:
			list = append(list, " bursts through the "+direction+" and breaks free for a massive gain of ",
				" passes the line of scrimmage, dodges right past a linebacker and makes a break for it before being taken down for ",
				" spins past a defender, runs toward the "+direction+" side and makes some headway! Bulldozes a defender but is quickly pulled down. Wow. Ran for ",
				" spins past a defender, runs toward the "+direction+" side. Bulldozes a defender but is quickly pulled down. Wow. Ran for ",
				" charges upfield, shrugging off tacklers and marking a significant run of ",
				" blazes down the field, eluding defenders and finally tackled after a fantastic run of ",
				" blazes down the field, eluding defenders and finally tackled. Fantastic gain of ",
				" gallops downfield like a freight train, eventually brought down after a tremendous run of ",
				" gallops downfield like a freight train, eventually brought down after a tremendous effort. A run for ")
		case yards > 9:
			list = append(list,
				" powers through the "+direction+" for a strong run, picking up ",
				" finds a hole in the "+direction+" and hits it with authority, chalking up ",
				" darts through the "+direction+" side, with a burst of speed, racking up ",
				" carves up the defense on the "+direction+" for a healthy gain of ",
				" exploits a gap to the "+direction+" and rumbles past defenders for ",
				" turns on the afterburners and leaves the defense in the dust, a solid run for ",
				" weaves through the defense to the "+direction+" and breaks away, tallying ",
				" finds daylight on the "+direction+", turning on the jets for a gain of ",
				" blasts past the line into the "+direction+" side, barreling forward for ")
		case yards > 4:
			list = append(list,
				" squeezes through a gap on the "+direction+" and grinds out ",
				" muscles his way past the line and into the "+direction+", securing ",
				" takes a handoff to the "+direction+" and barrels ahead for ",
				" follows his blockers to the "+direction+" and powers through for ",
				" with a hop and a skip, sidesteps a defender and sprints to the "+direction+" for ",
				" hits the hole on the "+direction+" with gusto, picking up ",
				" jukes to the "+direction+" and finds a seam, bolting for ",
				" with quick feet, darts to the "+direction+" and snags ")
		case yards > 0:
			list = append(list,
				" nudges forward on the "+direction+", managing to get ",
				" makes a run for it but is brought down quickly. Ran for ",
				" speeds past a defender but is swiftly tackled. Ran for ",
				" scoots up the field for ",
				" squeezes past the line of scrimmage but is quickly halted. A gain of ",
				" charges forward but meets a wall of defenders. A short gain of ",
				" finds little room and is stopped after a modest gain of ",
				" stumbles slightly but regains footing, a short return for ",
				" gets a grip on the ball and scurries ahead, only managing ",
				" gathers the kick and charges ahead into the arms of the opposition, a minimal gain of ")
		case yards == 0:
			list = append(list, " is met at the line of scrimmage, no gain on the play. ",
				" is immediately stopped at the line of scrimmage, no gain on the play. ",
				" immediately bumps into a lineman. No gain on the play. ",
				" is immediately tackled by the defensive line. No gain on the play. ")
		default: // Negative yardage
			list = append(list, " is wrapped up behind the line, a loss of ",
				" is stopped in the backfield for ",
				" meets a defender and is tackled in the backfield for ",
				" tries to find a gap, and just like that he's stopped for ",
				" finds no gap, tackled for ",
				" finds no gap, brought down for ",
				" attempts to find a gap and is brought down for ")
		}
	}
	return list
}

func getDrawRunList(yards int, isleft, touchdown, outOfBounds, twoPtConversion, fumble, safety bool) []string {
	list := []string{}
	direction := "left"
	if isleft {
		direction = "right"
	}
	// Depending on the yards gained or lost, append appropriate descriptions
	if touchdown {
		list = append(list, " deceives the defense with a draw and sprints into the end zone for a touchdown! ",
			" executes the draw play to perfection, finding a path to the end zone for six! ",
			" capitalizes on the draw, darting through the confused defenders for a touchdown! ")
	} else if outOfBounds {
		list = append(list, " takes the draw to the "+direction+" and steps out after gaining ",
			" takes the draw to the "+direction+" and steps out, gaining ",
			" takes the draw to the "+direction+" and steps out out bounds, gaining ",
			" takes the draw to the "+direction+" and steps to the sideline, gaining ",
			" uses the draw to find space and then smartly steps out of bounds after a decent gain of ")
	} else if twoPtConversion && touchdown {
		list = append(list, " fools the defense with a draw and secures the two-point conversion! ",
			" slips through the line on a draw play, converting for two points! ")
	} else if twoPtConversion && !touchdown {
		list = append(list, " attempts a draw for the two-point conversion but is stopped short! ",
			" can't find the end zone on the draw play, missing the conversion! ")
	} else if fumble {
		list = append(list, " fumbles the ball on the draw play after advancing ",
			" loses the ball on the draw, turning it over after a brief gain of ")
	} else if safety {
		list = append(list, " is trapped in the end zone on a draw play, resulting in a safety! ",
			" can't escape the clutches of the defense on the draw, conceding a safety! ")
	} else {
		switch {
		case yards > 14:
			list = append(list, " fools everyone with a draw and breaks free for a massive gain of ",
				" fools the defense with a draw and breaks free for a massive gain of ",
				" fools the defense with a draw, breaking free for a massive gain of ",
				" fools the defense with a draw, breaking free for a strong gain of ",
				" fools the defense with a draw, breaking free for a fantastic gain of ",
				" fools the defense with a draw, breaking free for ",
				" turns the draw play into a highlight, dashing past the defense for ",
				" turns the draw play into a highlight, sprinting past the defense for ",
				" turns the draw play into a highlight, sprinting past for ",
				" turns the draw play into a highlight, dashing past for ",
				" turns the draw play into a highlight, running past the defense for ",
				" turns the draw play into a highlight, running for ")
		case yards > 9:
			list = append(list, " navigates through the line on a draw, pushing ahead for ",
				" navigates through the line on a draw, gaining ",
				" navigates through the line on a draw, a push for ",
				" navigates through the line on a draw, a strong gain of ",
				" finds a seam on the draw and accelerates for a strong gain of ",
				" finds a seam on the draw and accelerates for a fantastic ",
				" finds a seam on the draw and accelerates for ",
				" finds a seam on the draw and runs for a gain of ",
				" finds a seam on the draw and runs for a strong gain of ",
				" finds a seam on the draw and makes a dash for it, gaining ")
		case yards > 4:
			list = append(list, " grinds out yards with a well-executed draw, managing to get ",
				" picks his way through the defense on the draw for a solid ",
				" picks his way through the defense on the draw for a gain of ",
				" picks his way through the defense on the draw gaining ",
				" finds a gap on the draw and makes a run for ",
				" finds a gap on the draw and makes a gain of ",
				" finds a gap on the draw and gains ",
				" finds a gap on the draw, gaining ")
		case yards > 0:
			list = append(list, " gains a few tough yards on the draw, every inch fought for ",
				" makes a small but positive gain on the draw, inching forward for ",
				" makes a run for it on the draw, gaining ")
		default: // Negative yardage
			list = append(list, " is swallowed up behind the line on the draw, losing ",
				" can't fool the defense on the draw, losing ",
				" can't find a gap on the draw and is tackled for a loss of ")
		}
	}
	return list
}

func getReadOptionList(yards int, isleft, touchdown, outOfBounds, twoPtConversion, fumble, safety bool) []string {
	list := []string{}
	direction := "left"
	if isleft {
		direction = "right"
	}
	// Depending on the yards gained or lost, append appropriate descriptions
	if touchdown {
		list = append(list, " makes the perfect read and takes it all the way to the house for a touchdown! ",
			" reads the defense, opts to keep it, and dashes to the end zone for six! ",
			" deceives the defense with a slick read option, sprinting into the end zone for a score! ")
	} else if outOfBounds {
		list = append(list, " opts for the "+direction+" and steps out after a decent gain. ",
			" reads, keeps, and sprints to the "+direction+", going out of bounds after gaining yards. ")
	} else if twoPtConversion && touchdown {
		list = append(list, " executes the read option flawlessly for a two-point conversion! ",
			" makes the right read and converts for two crucial points! ")
	} else if twoPtConversion && !touchdown {
		list = append(list, " tries the read option for the conversion but gets stopped short! ",
			" opts to keep but can't break through for the two points. ")
	} else if fumble {
		list = append(list, " fumbles while executing the read option, losing the ball after gaining yards. ",
			" makes a read but loses the ball on the run, a costly turnover. ")
	} else if safety {
		list = append(list, " gets trapped in the end zone on a read option, resulting in a safety! ",
			" opts to keep but is caught in the end zone for a safety! ")
	} else {
		switch {
		case yards > 14:
			list = append(list, " makes a great read and breaks free, galloping for a huge gain of ",
				" reads the defense perfectly and darts through the "+direction+" for a big pickup of ")
		case yards > 9:
			list = append(list, " opts to keep and finds a seam, powering ahead for ",
				" makes the right call on the read option, rushing for a solid gain of ")
		case yards > 4:
			list = append(list, " keeps on the read option and maneuvers for a decent pickup of ",
				" reads the play well, keeping it and grinding out ")
		case yards > 0:
			list = append(list, " decides to keep but is quickly brought down after gaining ",
				" makes a quick read and pushes forward for a short gain of ")
		case yards == 0:
			list = append(list, " is met at the line of scrimmage on the option, no gain on the play. ",
				" is immediately stopped at the line of scrimmage on the option, no gain on the play. ",
				" is immediately tackled by the defensive line. No gain on the play. ")
		default: // Negative yardage
			list = append(list, " makes a read but is swarmed behind the line, losing ",
				" opts to keep but is tackled for a loss of ")
		}
	}
	return list
}

// func getSpeedOptionList(yards int, isleft, touchdown, outOfBounds, twoPtConversion, fumble, safety bool) []string {
// 	list := []string{}
// 	direction := "left"
// 	if isleft {
// 		direction = "right"
// 	}
// 	// Depending on the yards gained or lost, append appropriate descriptions
// 	if touchdown {
// 		list = append(list, "")
// 	} else if outOfBounds {
// 		list = append(list, "")
// 	} else if twoPtConversion && touchdown {
// 		list = append(list, "")
// 	} else if twoPtConversion && !touchdown {
// 		list = append(list, "")
// 	} else if fumble {
// 		list = append(list, "")
// 	} else if safety {
// 		list = append(list, "")
// 	} else {
// 		switch {
// 		case yards > 14:
// 			list = append(list, "")
// 		case yards > 9:
// 			list = append(list, "")
// 		case yards > 4:
// 			list = append(list, "")
// 		case yards > 0:
// 			list = append(list, "")
// 		case yards == 0:
// 			list = append(list, " is met at the line of scrimmage, no gain on the play. ",
// 				" is immediately stopped at the line of scrimmage, no gain on the play. ",
// 				" is immediately tackled by the defensive line. No gain on the play. ")
// 		default: // Negative yardage
// 			list = append(list, "")
// 		}
// 	}
// 	return list
// }

// func getInvertedOptionList(yards int, isleft, touchdown, outOfBounds, twoPtConversion, fumble, safety bool) []string {
// 	list := []string{}
// 	direction := "left"
// 	if isleft {
// 		direction = "right"
// 	}
// 	// Depending on the yards gained or lost, append appropriate descriptions
// 	if touchdown {
// 		list = append(list, "")
// 	} else if outOfBounds {
// 		list = append(list, "")
// 	} else if twoPtConversion && touchdown {
// 		list = append(list, "")
// 	} else if twoPtConversion && !touchdown {
// 		list = append(list, "")
// 	} else if fumble {
// 		list = append(list, "")
// 	} else if safety {
// 		list = append(list, "")
// 	} else {
// 		switch {
// 		case yards > 14:
// 			list = append(list, "")
// 		case yards > 9:
// 			list = append(list, "")
// 		case yards > 4:
// 			list = append(list, "")
// 		case yards > 0:
// 			list = append(list, "")
// 		case yards == 0:
// 			list = append(list, " is met at the line of scrimmage, no gain on the play. ",
// 				" is immediately stopped at the line of scrimmage, no gain on the play. ",
// 				" is immediately tackled by the defensive line. No gain on the play. ")
// 		default: // Negative yardage
// 			list = append(list, "")
// 		}
// 	}
// 	return list
// }

// func getTripleOptionList(yards int, isleft, touchdown, outOfBounds, twoPtConversion, fumble, safety bool) []string {
// 	list := []string{}
// 	direction := "left"
// 	if isleft {
// 		direction = "right"
// 	}
// 	// Depending on the yards gained or lost, append appropriate descriptions
// 	if touchdown {
// 		list = append(list, "")
// 	} else if outOfBounds {
// 		list = append(list, "")
// 	} else if twoPtConversion && touchdown {
// 		list = append(list, "")
// 	} else if twoPtConversion && !touchdown {
// 		list = append(list, "")
// 	} else if fumble {
// 		list = append(list, "")
// 	} else if safety {
// 		list = append(list, "")
// 	} else {
// 		switch {
// 		case yards > 14:
// 			list = append(list, "")
// 		case yards > 9:
// 			list = append(list, "")
// 		case yards > 4:
// 			list = append(list, "")
// 		case yards > 0:
// 			list = append(list, "")
// 		case yards == 0:
// 			list = append(list, " is met at the line of scrimmage, no gain on the play. ",
// 				" is immediately stopped at the line of scrimmage, no gain on the play. ",
// 				" is immediately tackled by the defensive line. No gain on the play. ")
// 		default: // Negative yardage
// 			list = append(list, "")
// 		}
// 	}
// 	return list
// }

/*
15: "Triple Option Left",
16: "Triple Option Right",
17: "Choice Inside Right",
18: "Choice Inside Left",
19: "Choice Outside Right",
20: "Choice Outside Left",
21: "Choice Power Right",
22: "Choice Power Left",
23: "Peek Inside Right",
24: "Peek Inside Left",
25: "Peek Outside Right",
26: "Peek Outside Left",
27: "Peek Power Right",
28: "Peek Power Left",
*/
