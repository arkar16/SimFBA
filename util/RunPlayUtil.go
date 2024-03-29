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
		list = append(list, getDrawRunList(yards, playName == "Draw Left", touchdown, outOfBounds, twoPtConversion, fumble, safety)...)
	} else if playName == "Read Option Left" || playName == "Read Option Right" {
		list = append(list, getReadOptionList(yards, playName == "Read Option Left", touchdown, outOfBounds, twoPtConversion, fumble, safety)...)
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
			" reads the goalline defense, opts to keep it, and dashes to the end zone for six! ",
			" deceives the goalline defense with a slick read option, sprinting into the end zone for a score! ",
			" navigates the read option with precision, blazing his way to the end zone for a spectacular touchdown! ",
			" outwits the goalline defense with a masterful read, darting to paydirt for the touchdown! ",
			" showcases his athleticism on the goalline, turning the read option into a thrilling touchdown run! ")
	} else if outOfBounds {
		list = append(list, " opts for the "+direction+" and steps out after a decent gain. ",
			" reads, keeps, and sprints to the "+direction+", going out of bounds after gaining yards. ",
			" leverages the read option to gain the edge and smartly steps out, banking ",
			" exploits the defense's hesitation on the read and scoots out of bounds for a gain of ",
			" exploits the defense's hesitation on the option and scoots out of bounds for a gain of ",
			" surprises the defensive line on the option and edges out of bounds for a gain of ",
			" surprises the defensive line on the option and edges out of bounds, gaining ")
	} else if twoPtConversion && touchdown {
		list = append(list, " executes the read option flawlessly on the goalline for a two-point conversion! ",
			" makes the right read in the red zone and converts for two crucial points! ",
			" brilliantly navigates the read option, crossing into the end zone for the extra points! ",
			" capitalizes on the read option, weaving his way for a successful two-point play! ")
	} else if twoPtConversion && !touchdown {
		list = append(list, " tries the read option for the conversion but gets stopped short! ",
			" opts to keep but can't break through for the two points. ",
			" hesitates on the read option, can't break through for the two points. ",
			" gives the read option a go for the two points, but the defense stands tall, denying the conversion. ",
			" opts for the read on the conversion attempt, yet the defense snuffs it out, thwarting the extra points. ",
			" stalls on the read option attempt and stopped short of the goalline, thwarting the extra points. ",
			" stalls on the read option attempt and stopped short of the end zone, thwarting the extra points. ",
			" stalls on the read option attempt and stopped short of the end zone, thwarting the two point attempt. ",
			" trips up on the read option attempt and stopped short of the goalline, thwarting the extra points. ",
			" trips up on the read option attempt and stopped short of the end zone, thwarting the extra points. ",
			" trips up on the read option attempt and stopped short of the end zone, thwarting the two point attempt. ",
			" attempts the read option and is stopped short of the goalline, thwarting the extra points. ",
			" attempts the read option and is stopped short of the end zone, thwarting the extra points. ",
			" attempts the read option and is stopped short of the end zone, thwarting the two point attempt. ")
	} else if fumble {
		list = append(list, " fumbles while executing the read option, losing the ball after gaining ",
			" makes a read but loses the ball on the run, a costly turnover after ",
			" tries to navigate the read option but ends up turning the ball over with a fumble after ",
			" keeps on the read option, only to lose control of the ball, squandering ")
	} else if safety {
		list = append(list, " gets trapped in the end zone on a read option, resulting in a safety! ",
			" opts to keep but is caught in the end zone for a safety! ",
			" disastrously misreads the option in the end zone, leading to a safety against his team. ",
			" hesitates on the option in the end zone, leading to a safety against his team. ",
			" hesitates on the read option in the end zone, leading to a safety against his team. ",
			" is enveloped by the defense in the end zone during the read option, resulting in a safety. ")
	} else {
		switch {
		case yards > 14:
			list = append(list, " makes a great read and breaks free, galloping for a huge gain of ",
				" reads the defense perfectly and darts through the "+direction+" for a big pickup of ",
				" takes the read option and finds a gap in the defense for a pickup of ",
				" takes advantage of the defense's anticipation on the read, breaking free for ",
				" surprises the defense on the read option and nimbly navigates the middle of the field for a big gain of ",
				" surprises the defense on the read option and nimbly navigates the "+direction+" side of the field for a big gain of ",
				" surprises the defense on the read option and elusively navigates the "+direction+" side of the field for a huge gain of ",
				" reads the defense on the option and elusively navigates the "+direction+" side of the field for a gain of ",
				" reads the defense on the option and hits the gas towards the "+direction+" side of the field for a huge gain of ",
				" reads the defense on the option and dashes towards the "+direction+" edge for a huge gain of ")
		case yards > 9:
			list = append(list, " opts to keep and finds a seam, powering ahead for ",
				" opts to keep and finds a gap on the "+direction+" side, powering ahead for ",
				" opts to keep and finds a seam on the "+direction+" side, powering ahead for ",
				" executes the option and finds a seam on the "+direction+" side, powering through for ",
				" executes the option and finds a gap on the "+direction+" side, powering through for ",
				" executes the option and finds a seam on the "+direction+" side, powering ahead for ",
				" executes the option and finds a seam on the "+direction+" side for a gain of ",
				" executes the option and finds a seam on the "+direction+" side gaining ",
				" executes the option and finds a seam on the "+direction+" side, a gain of ",
				" makes the right call on the read option, rushing for a solid gain of ")
		case yards > 4:
			list = append(list, " keeps on the read option and maneuvers for a decent pickup of ",
				" reads the play well, keeping it and grinding out ",
				" reads the play well, keeping the option and grinding out ",
				" reads the play well, keeping it and churning out ",
				" reads the play well, executes the option and running for ",
				" reads the play well, executes the option and running for a gain of ",
				" reads the play well, executes the option and running for a moderate gain of ",
				" surprises the defensive line on the option but is stopped shortly on the "+direction+" side for a moderate gain of ",
				" finds a gap on the defensive line on the option but is stopped shortly on the "+direction+" side for a gain of ",
				" finds a seam on the defensive line on the option but is stopped shortly on the "+direction+" side for a gain of ",
				" finds a seam through the read option but is stopped shortly on the "+direction+" side for a moderate gain of ")
		case yards > 0:
			list = append(list, " decides to keep but is quickly brought down after gaining ",
				" makes a quick read and pushes forward for a short gain of ",
				" makes a quick read and is brought down quickly a short gain of ",
				" makes a quick read and is brought down for short gain of ",
				" makes a quick read and is brought down shortly, gaining ")
		case yards == 0:
			list = append(list, " is met at the line of scrimmage on the option, no gain on the play. ",
				" is immediately stopped at the line of scrimmage on the option, no gain on the play. ",
				" attempts to execute the read option but cannot trick the defense, no gain on the play. ",
				" attempts to execute the read option but is stopped swiftly, no gain on the play. ",
				" attempts to execute the read option but is stopped on the line of scrimmage, no gain on the play. ",
				" attempts to execute the read option but is shut down on the line of scrimmage, no gain on the play. ",
				" attempts the read option but is shut down on the line of scrimmage, no gain on the play. ",
				" attempts the read and is immediately tackled on the line. No gain on the play. ",
				" is immediately tackled by the defensive line. No gain on the play. ")
		default: // Negative yardage
			list = append(list,
				" makes a read but is swarmed behind the line, losing ",
				" makes a read but is met before the line of scrimmage, losing ",
				" makes a read but is tackled behind the line, losing ",
				" makes a read but is stopped behind the line, losing ",
				" makes a read but is stopped swiftly, losing ",
				" attempts a read but is swarmed behind the line, losing ",
				" attempts a read but is met before the line of scrimmage, losing ",
				" attempts a read but is tackled behind the line, losing ",
				" attempts a read but is stopped behind the line, losing ",
				" attempts a read but is stopped swiftly, losing ",
				" opts to keep but is tackled for a loss of ",
			)
		}
	}
	return list
}

func getSpeedOptionList(yards int, isleft, touchdown, outOfBounds, twoPtConversion, fumble, safety bool) []string {
	list := []string{}
	direction := "left"
	if isleft {
		direction = "right"
	}
	// Depending on the yards gained or lost, append appropriate descriptions
	if touchdown {
		list = append(list,
			" executes the speed option flawlessly, pitching it to the running back who dashes to the "+direction+" and into the end zone!",
			" perfects the pitch on the speed option, leading to a breathtaking dash down the "+direction+" sideline for a touchdown!",
			" orchestrates the speed option to perfection, resulting in an electrifying touchdown sprint to the "+direction+"!")
	} else if outOfBounds {
		list = append(list,
			" pitches it out on the speed option and the runner bolts to the "+direction+", stepping out after a solid gain of ",
			" utilizes the speed option, sending the back towards the "+direction+" where he's pushed out after a decent advance ")
	} else if twoPtConversion && touchdown {
		list = append(list,
			" nails the speed option for a successful two-point conversion, with the runner breaking to the "+direction+" and crossing the plane! ",
			" capitalizes on the speed option, pitching it for a successful dash into the end zone for two points! ")
	} else if twoPtConversion && !touchdown {
		list = append(list,
			" attempts the speed option for two points but the defense clamps down, thwarting the conversion to the "+direction+". ",
			" tries to convert with a speed option but the pitch to the "+direction+" is stopped, no extra points. ")
	} else if fumble {
		list = append(list,
			" makes a run for it on the speed option but loses control of the ball after gaining ",
			" makes a break for it on the speed option and he's lost the ball! It's a fumble! What a turn of events after ",
			" makes a break for it on the speed option and with the tackl- it's a fumble! What a turn of events after ",
			" makes a break for it on the speed option and with the tack- it's a fumble! What a turn of events after ",
			" makes a run for it on the option and loses the football on the tackle! ",
			" makes a run for it on the option and loses the ball after being brought down for ",
			" fumbles the pitch on the speed option, losing control of the ball after a minor gain of ",
			" fumbles the pitch on the speed option and loses the ball after a minor gain of ",
			" sees the speed option go awry as the ball hits the ground during the pitch, missing ",
			" mistakes the pitch and loses control of the ball, fumbling for ")
	} else if safety {
		list = append(list,
			" disastrously opts for a speed option that ends in a safety, with the runner tackled in the end zone! ",
			" tries a risky speed option in the end zone, which backfires for a safety against them! ",
			"attempts the speed option but cannot fool the defense, and that's a safety! ",
			"attempts the speed option but can't fool the defense, and that's a safety! ")
	} else {
		switch {
		case yards > 14:
			list = append(list,
				" tricks the defense on the speed option and tears down the "+direction+" for a huge gain of ",
				" tricks the defense on the speed option and runs to the "+direction+" for a huge gain of ",
				" surprises the defense on the speed option and runs to the "+direction+" for a huge gain of ",
				" surprises the defense on the speed option and tears down the "+direction+" for a huge gain of ",
				" surprises the defense on the speed option and runs to the "+direction+" for a gain of ",
				" surprises the defense on the speed option and runs to the "+direction+" for a huge gain of ",
				" masterfully executes the speed option and tears down the "+direction+" for a huge gain of ",
				" masterfully executes the speed option and runs to the "+direction+" for a gain of ",
				" masterfully executes the speed option and runs to the "+direction+" for a huge gain of ",
				" breaks a tackle on the speed option and makes a run for the "+direction+" for a huge pickup of ",
				" breaks a tackle on the speed option and makes a break for the "+direction+" side for a huge pickup of ",
				" breaks a tackle on the speed option and dashes for the "+direction+" sideline for a huge pickup of ",
				" breaks a tackle on the speed option and dashes downfield towards the "+direction+" sideline for a huge pickup of ",
				" breaks a tackle on the speed option and breaks downfield towards the "+direction+" sideline for a huge pickup of ",
				" skillfully executes a speed option, leading to a sprint along the "+direction+" that racks up significant yardage of ")
		case yards > 9:
			list = append(list,
				" opts for the speed option, making a timely pitch that results in a brisk "+direction+" run for ",
				" opts for the speed option, succeeding the pitch that results in a brisk "+direction+" run for ",
				" opts for the speed option, succeeding the pitch that results in a brisk "+direction+" run for a large gain of ",
				" opts for the speed option, succeeding the pitch that results in a brisk "+direction+" run for a huge gain of ",
				" opts for the speed option, making a timely pitch that results in a swift "+direction+" run for ",
				" opts for the speed option, succeeding the pitch that results in a swift "+direction+" run for ",
				" opts for the speed option, making a timely pitch that results in a swift "+direction+" run for a large gain of ",
				" opts for the speed option, succeeding the pitch that results in a swift "+direction+" run for a large gain of ",
				" eludes defenders on the speed option and makes some ground before being brought down for ",
				" eludes defenders on the speed option and makes some ground before being brought down for a gain of ",
				" eludes defenders on the speed option and darts to the "+direction+" before being brought down for a gain of ",
				" eludes defenders on the speed option and darts to the "+direction+" before being brought down for a solid gain of ",
				" eludes defenders on the speed option and darts to the "+direction+" before being brought down, gaining ",
				" eludes defenders on the speed option and runs to the "+direction+" before being brought down for a gain of ",
				" eludes defenders on the speed option and runs to the "+direction+" before being brought down for a solid gain of ",
				" eludes defenders on the speed option and runs to the "+direction+" before being brought down, gaining ",
				" misses a tackle on the speed option and runs to the "+direction+" before being brought down for a gain of ",
				" misses a tackle on the speed option and runs to the "+direction+" before being brought down for a solid gain of ",
				" misses a tackle on the speed option and runs to the "+direction+" before being brought down, gaining ",
				" capitalizes on the speed option, breaking to the "+direction+" for a solid pickup of ",
				" capitalizes on the option, breaking to the "+direction+" for a solid pickup of ",
				" capitalizes on the speed option, breaking to the "+direction+" for a solid gain of ",
				" capitalizes on the option, breaking to the "+direction+" for a solid gain of ",
				" capitalizes on the speed option, breaking to the "+direction+" for a solid run of ",
				" capitalizes on the option, breaking to the "+direction+" for a run of ")
		case yards > 4:
			list = append(list,
				" carries out on the speed option to the "+direction+" for a moderate gain of ",
				" carries out on the speed option to the "+direction+" for a gain of ",
				" carries out on the speed option to the "+direction+", gaining ",
				" goes with the speed option, going around the "+direction+" edge for a respectable gain of ",
				" goes with the speed option, going around the "+direction+" edge for a gain of ",
				" goes with the speed option, going around the "+direction+" edge, gaining ",
				" takes the speed option to the "+direction+" edge for a respectable gain of ",
				" takes the speed option to the "+direction+" edge for a gain of ",
				" takes the speed option to the "+direction+" edge, gaining ",
				" takes the speed option to the "+direction+" and makes some ground for a moderate gain of ",
				" takes the speed option to the "+direction+" and makes some ground for a gain of ",
				" takes the speed option to the "+direction+" and makes some ground, gaining",
			)
		case yards > 0:
			list = append(list,
				" manages to eke out a few yards on a speed option to the "+direction+", advancing for ",
				" manages to eke out a few yards on the option pitch to the "+direction+", advancing for ",
				" manages to eke out a few yards on the option pitch to the "+direction+", a small gain of ",
				" manages to eke out a few yards on the option pitch to the "+direction+", gaining ",
				" edges out a few yards on the option pitch to the "+direction+", advancing for ",
				" edges out a few yards on the option pitch to the "+direction+", a small gain of ",
				" edges out a few yards on the option pitch to the "+direction+", gaining ",
				" takes the option to the "+direction+" side and is brought down, advancing for ",
				" takes the option to the "+direction+" side and is brought down, a small gain of ",
				" takes the option to the "+direction+" side and is brought down, gaining ",
				" takes the option to the "+direction+" side and is tackled, advancing for ",
				" takes the option to the "+direction+" side and is tackled, a small gain of ",
				" takes the option to the "+direction+" side and is tackled, gaining ",
				" sees a modest return on the speed option, with the pitch leading to a slight gain of ")
		case yards == 0:
			list = append(list, " is met at the line of scrimmage, no gain on the play. ",
				" is immediately stopped at the line of scrimmage, no gain on the play. ",
				" is immediately tackled by the defensive line. No gain on the play. ")
		default: // Negative yardage
			list = append(list,
				" is stopped on the speed option as two defenders spy both option players. Caught behind the line for a loss of ",
				" is stopped on the speed option as two defenders spy both option players, losing ",
				" has trouble with the option pitch, handles the ball and goes down behind the line, losing ",
				" has trouble with the option pitch and goes down behind the line, losing ",
				" has trouble with the option pitch and goes down behind the line for a loss of ",
				" hesitates on the speed option pitch and goes down behind the line for a loss of ",
				" hesitates on the speed option pitch and is tackled behind the line for a loss of ",
				" hesitates on the speed option pitch and is tackled behind the line, losing ",
				" hesitates on the speed option pitch and is stopped behind the line, losing ",
				" hesitates on the speed option pitch and is stopped behind the line for a loss of",
				" can't escape the edge on the option and is stopped behind the line for a loss of",
				" can't escape the edge on the option and is stopped behind the line, losing",
				" misreads the speed option and is stopped behind the line, losing",
				" misreads the speed option and is tackled behind the line, losing",
				" misreads the speed option and goes down behind the line, losing",
				" misreads the speed option and is manhandled behind the line, losing",
				" misreads the speed option and is manhandled behind the line for a loss of",
				" sees the speed option falter, with the runner caught behind the line for a loss of ",
				" watches the speed option play crumble, resulting in negative yardage of ")
		}
	}
	return list
}

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
