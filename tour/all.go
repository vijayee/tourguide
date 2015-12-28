package tour

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

func Init() {
	for _, t := range allTopics {
		Topics[t.ID] = t
		IDs = append(IDs, t.ID)
	}

	sort.Sort(IDSlice(IDs))
}

// TODO move content into individual files if desired

// TODO(brian): If sub-topics are needed, write recursively (as tree comprised
// of Section nodes:
//
// type Section interface {
// 	Sections() []Section
// 	Topic() Topic
// }

var (
	// TODO bootstrapping

	// TODO pinning: ensuring a block is kept in local storage (i.e. not
	// evicted from cache).

	Introduction = Chapter(0)
	FileBasics   = Chapter(1)
	NodeBasics   = Chapter(2)
	MerkleDag    = Chapter(3)
	Network      = Chapter(4)
	Daemon       = Chapter(5)
	Routing      = Chapter(6)
	Exchange     = Chapter(7)
	Ipns         = Chapter(8)
	Mounting     = Chapter(9)
	Plumbing     = Chapter(10)
	Formats      = Chapter(11)
)

// Topics contains a mapping of Tour Topic ID to Topic
var allTopics = []Topic{
	Topic{ID: Introduction(0), Content: IntroHelloMars, hasPassed: false},
	Topic{ID: Introduction(1), Content: IntroTour, hasPassed: false},
	Topic{ID: Introduction(2), Content: IntroAboutIpfs, hasPassed: false},

	Topic{ID: FileBasics(1), Content: FileBasicsFilesystem, hasPassed: false},
	Topic{ID: FileBasics(2), Content: FileBasicsGetting, hasPassed: false},
	Topic{ID: FileBasics(3), Content: FileBasicsAdding, hasPassed: false},
	Topic{ID: FileBasics(4), Content: FileBasicsDirectories, hasPassed: false},
	Topic{ID: FileBasics(5), Content: FileBasicsDistributed, hasPassed: false},
	Topic{ID: FileBasics(6), Content: FileBasicsMounting, hasPassed: false},

	Topic{NodeBasics(0), NodeBasicsInit, false},
	Topic{NodeBasics(1), NodeBasicsHelp, false},
	Topic{NodeBasics(2), NodeBasicsUpdate, false},
	Topic{NodeBasics(3), NodeBasicsConfig, false},

	Topic{MerkleDag(0), MerkleDagIntro, false},
	Topic{MerkleDag(1), MerkleDagContentAddressing, false},
	Topic{MerkleDag(2), MerkleDagContentAddressingLinks, false},
	Topic{MerkleDag(3), MerkleDagRedux, false},
	Topic{MerkleDag(4), MerkleDagIpfsObjects, false},
	Topic{MerkleDag(5), MerkleDagIpfsPaths, false},
	Topic{MerkleDag(6), MerkleDagImmutability, false},
	Topic{MerkleDag(7), MerkleDagUseCaseUnixFS, false},
	Topic{MerkleDag(8), MerkleDagUseCaseGitObjects, false},
	Topic{MerkleDag(9), MerkleDagUseCaseOperationalTransforms, false},

	Topic{Network(0), Network_Intro, false},
	Topic{Network(1), Network_Ipfs_Peers, false},
	Topic{Network(2), Network_Daemon, false},
	Topic{Network(3), Network_Routing, false},
	Topic{Network(4), Network_Exchange, false},
	Topic{Network(5), Network_Intro, false},

	// TODO daemon - {API, API Clients, Example} how old-school http + ftp
	// clients show it
	Topic{Daemon(0), Daemon_Intro, false},
	Topic{Daemon(1), Daemon_Running_Commands, false},
	Topic{Daemon(2), Daemon_Web_UI, false},

	Topic{Routing(0), Routing_Intro, false},
	Topic{Routing(1), Rouing_Interface, false},
	Topic{Routing(2), Routing_Resolving, false},
	Topic{Routing(3), Routing_DHT, false},
	Topic{Routing(4), Routing_Other, false},

	// TODO Exchange_Providing
	// TODO Exchange_Providers
	Topic{Exchange(0), Exchange_Intro, false},
	Topic{Exchange(1), Exchange_Getting_Blocks, false},
	Topic{Exchange(2), Exchange_Strategies, false},
	Topic{Exchange(3), Exchange_Bitswap, false},

	Topic{Ipns(0), Ipns_Name_System, false},
	Topic{Ipns(1), Ipns_Mutability, false},
	Topic{Ipns(2), Ipns_PKI_Review, false},
	Topic{Ipns(3), Ipns_Publishing, false},
	Topic{Ipns(4), Ipns_Resolving, false},
	Topic{Ipns(5), Ipns_Consistency, false},
	Topic{Ipns(6), Ipns_Records_Etc, false},

	Topic{Mounting(0), Mounting_General, false},
	Topic{Mounting(1), Mounting_Ipfs, false},
	Topic{Mounting(2), Mounting_Ipns, false},

	Topic{Plumbing(0), Plumbing_Intro, false},
	Topic{Plumbing(1), Plumbing_Ipfs_Block, false},
	Topic{Plumbing(2), Plumbing_Ipfs_Object, false},
	Topic{Plumbing(3), Plumbing_Ipfs_Refs, false},
	Topic{Plumbing(4), Plumbing_Ipfs_Ping, false},
	Topic{Plumbing(5), Plumbing_Ipfs_Id, false},

	Topic{Formats(0), Formats_MerkleDag, false},
	Topic{Formats(1), Formats_Multihash, false},
	Topic{Formats(2), Formats_Multiaddr, false},
	Topic{Formats(3), Formats_Multicodec, false},
	Topic{Formats(4), Formats_Multicodec, false},
	Topic{Formats(5), Formats_Multikey, false},
	Topic{Formats(6), Formats_Protocol_Specific, false},
}

// Introduction

var IntroHelloMars = Content{
	Title: "Hello Mars",
	Text: `Chapter 1: Hello Mars 
So long ago.......In a galaxy far far away...Unixians created life and man out of very simple pieces that fit together like lego bricks. They then placed them on the planet earth as stewards and care takers of the primitive planet. Their leader began to give very specific instructions that boomed from the sky to man on the nature of creation, the storage of knowledge, and the Universe which stated "Behold I have created you out of tiny interlocking pieces which store the potential and capacity for great knowledge and attainment.....and....!@#%@$@$$" Their leader looked to the heavens in the direction of mars and became distraught and angry whisper echoed the earth. "Why is Mars still red!?....What do you mean we spent our water budget on earth!?.....how long till development can continue!?!?!.....It was your job to manage the budget!!....I'm mid press-release what will I tell these guys!?!". Their leader then addressed man and said "Hold on, I'll be right back". Man was left confused and alone which soon led to great anger and frustration. It became customary to utter "!@#%@$@$$" in fluent unixian when angry which led to war. The Unixians returned moments later relative to their time but failed to account for relativity in time and discovered man had formed its own civilizations and organizational structures and a network of information called the internet. It was close in concept to the nature of knowledge they had tried to initially bestow upon them but all wrong in its structure and mostly composed of cats. So over-budget, the Unixians could not afford to start over and began to inspire those with the vision to fix the errors in their network of information. One robust enough to reach the stars. The Interplanetary File System began its creation in order to fix the problems created in ignorance. In order to receive the information you must issue commands to initialize with the heavens and your peers. You may use the primitive earth pipe to confirm your commands. ("Warp pipes are on the way!" -The Prophet Mario )

ipfs init | ./tourguide
	`,
	verify: func(stdin []byte) (bool, error) {
		return true, nil
	},
}
var IntroTour = Content{
	Title: "Hello Mars",
	Text: `
	Bacon ipsum dolor amet shankle drumstick hamburger capicola. Hamburger cow jerky meatloaf short loin capicola filet mignon shank shankle. Swine turducken jowl ground round landjaeger filet mignon pastrami alcatra. Bresaola tri-tip pastrami hamburger shankle t-bone chuck ham fatback flank alcatra tenderloin doner pork chop prosciutto. Turkey beef ribs ribeye tenderloin shankle, pastrami flank sausage frankfurter leberkas bresaola porchetta. Ham bresaola cow pork belly. Bresaola shank ham, tail beef ribs tenderloin cow.

Cupim beef shank, bacon t-bone swine pancetta hamburger brisket ham cow. Drumstick kevin ham, ribeye doner kielbasa shankle spare ribs brisket pancetta sausage ball tip. Spare ribs landjaeger biltong leberkas pork loin tongue kielbasa jowl ham pastrami chuck doner pig short ribs. Bacon biltong pig, pork kevin short ribs venison. Hamburger swine biltong tri-tip ball tip, pork belly jowl picanha bresaola beef pig tongue meatball.

Kevin biltong venison pastrami flank fatback. Beef ribs chuck kielbasa alcatra shankle short loin jerky, leberkas spare ribs brisket cow salami pastrami tail. Pancetta turkey beef ribs tenderloin chicken, venison bacon flank ham hock andouille shank kevin. Ground round boudin pancetta kielbasa capicola strip steak tenderloin turkey short loin.

Bresaola picanha cupim ground round, beef tri-tip jerky strip steak capicola t-bone. Shoulder prosciutto beef ribs, chuck sausage venison biltong. Sausage shoulder pig hamburger beef ribs short ribs shankle bresaola salami sirloin bacon flank swine. Picanha ball tip pork loin jerky kevin pork short loin meatloaf. Pork belly tongue shankle drumstick rump bresaola frankfurter ground round tail pork chop doner fatback cupim landjaeger kielbasa. Ball tip corned beef chuck pastrami. Pancetta swine leberkas bacon ribeye spare ribs.

Meatball ham kevin, cow venison fatback salami flank leberkas sausage beef ribs andouille shank tongue. Pastrami beef ribs biltong pork loin strip steak, drumstick ham hock picanha ham pork belly. Shankle flank t-bone, corned beef picanha chuck venison. Leberkas chicken corned beef pig salami boudin.

Does your lorem ipsum text long for something a little meatier? Give our generator a try… it’s tasty!
	`,
	verify: func(stdin []byte) (bool, error) {
		var input = string(stdin)
		if strings.Contains(input, "initializing ipfs node at") {
			return true, nil
		} else {
			return false, errors.New("Initialization Failed")
		}

	},
}
var IntroAboutIpfs = Content{
	Title: "About IPFS",
	verify: func(stdin []byte) (bool, error) {
		var input = string(stdin)
		if strings.Contains(input, "initializing ipfs node at") {
			return true, nil
		} else {
			return false, errors.New("Initialization Failed")
		}

	},
}

// File Basics

var FileBasicsFilesystem = Content{
	Title: "Filesystem",
	Text: `
	`,
}
var FileBasicsGetting = Content{
	Title: "Getting Files",
	Text: `ipfs cat
	`,
}
var FileBasicsAdding = Content{
	Title: "Adding Files",
	Text: `ipfs add
	`,
	verify: func(stdin []byte) (bool, error) {
		var input = string(stdin)
		re, err := regexp.Compile("added\\s([A-Za-z0-9]*)\\s([A-Za-z0-9]*)")
		if err != nil {
			return false, err
		}
		matches := re.FindStringSubmatch(input)
		if len(matches) < 1 {
			return false, errors.New("Add function failed:\n" + input)
		}
		if matches[1] != "" {
			return true, nil
		}
		return false, errors.New("Verification failed")

	},
}
var FileBasicsDirectories = Content{
	Title: "Directories",
	Text: `ipfs ls
	`,
	verify: func(stdin []byte) (bool, error) {
		var input = string(stdin)
		re, err := regexp.Compile("([A-Za-z0-9]*)\\s([0-9]*)\\s")
		if err != nil {
			return false, err
		}
		matches := re.FindStringSubmatch(input)
		if len(matches) < 1 {
			return false, errors.New("ls function failed:\n" + input)
		}
		if matches[1] != "" {
			return true, nil
		}
		return false, errors.New("Verification failed")
	},
}
var FileBasicsDistributed = Content{
	Title: "Distributed",
	Text: `ipfs cat from mars
	`,
}
var FileBasicsMounting = Content{
	Title: "Getting Files",
	Text: `ipfs mount (simple)
	`,
}

// Node Basics

var NodeBasicsInit = Content{
	Title: "Basics - init",

	// TODO touch on PKI
	//
	// This is somewhat relevant at ipfs init since the generated key pair is the
	// basis for the node's identity in the network. A cursory nod may be
	// sufficient at that stage, and goes a long way in explaining init's raison
	// d'être.
	// NB: user is introduced to ipfs init before ipfs add.
	Text: `
	`,
}
var NodeBasicsHelp = Content{
	Title: "Basics - help",
	Text: `
	`,
}
var NodeBasicsUpdate = Content{
	Title: "Basics - update",
	Text: `
	`,
}
var NodeBasicsConfig = Content{
	Title: "Basics - config",
	Text: `
	`,
}

// Merkle DAG
var MerkleDagIntro = Content{}
var MerkleDagContentAddressing = Content{}
var MerkleDagContentAddressingLinks = Content{}
var MerkleDagRedux = Content{}
var MerkleDagIpfsObjects = Content{}
var MerkleDagIpfsPaths = Content{}
var MerkleDagImmutability = Content{
	Title: "Immutability",
	Text: `
	TODO plan9
	TODO git
	`,
}

var MerkleDagUseCaseUnixFS = Content{}
var MerkleDagUseCaseGitObjects = Content{}
var MerkleDagUseCaseOperationalTransforms = Content{}

var Network_Intro = Content{}
var Network_Ipfs_Peers = Content{}
var Network_Daemon = Content{}
var Network_Routing = Content{}
var Network_Exchange = Content{}
var Network_Naming = Content{}

var Daemon_Intro = Content{}
var Daemon_Running_Commands = Content{}
var Daemon_Web_UI = Content{}

var Routing_Intro = Content{}
var Rouing_Interface = Content{}
var Routing_Resolving = Content{}
var Routing_DHT = Content{}
var Routing_Other = Content{}

var Exchange_Intro = Content{}
var Exchange_Bitswap = Content{}
var Exchange_Strategies = Content{}
var Exchange_Getting_Blocks = Content{}

var Ipns_Consistency = Content{}
var Ipns_Mutability = Content{}
var Ipns_Name_System = Content{}
var Ipns_PKI_Review = Content{
	Title: "PKI Review",
	Text: `
	TODO sign verify
	`,
}
var Ipns_Publishing = Content{}
var Ipns_Records_Etc = Content{}
var Ipns_Resolving = Content{}

var Mounting_General = Content{} // TODO note fuse
var Mounting_Ipfs = Content{}    // TODO cd, ls, cat
var Mounting_Ipns = Content{}    // TODO editing

var Plumbing_Intro = Content{}
var Plumbing_Ipfs_Block = Content{}
var Plumbing_Ipfs_Object = Content{}
var Plumbing_Ipfs_Refs = Content{}
var Plumbing_Ipfs_Ping = Content{}
var Plumbing_Ipfs_Id = Content{}

var Formats_MerkleDag = Content{}
var Formats_Multihash = Content{}
var Formats_Multiaddr = Content{}
var Formats_Multicodec = Content{}
var Formats_Multikey = Content{}
var Formats_Protocol_Specific = Content{}
