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
	Text: `
	Bacon ipsum dolor amet swine biltong pork loin tail shoulder short loin. Filet mignon spare ribs chuck, kevin ribeye tail pancetta. Biltong salami landjaeger jowl. Ham turkey biltong, swine kielbasa alcatra doner shank rump picanha chuck.

Short loin frankfurter prosciutto tail, bresaola boudin flank picanha ham hock kevin. Beef turducken pork belly cupim. Capicola alcatra prosciutto strip steak brisket turkey chuck shank jerky picanha ground round. Turkey sirloin pork loin picanha t-bone ribeye.
Bacon ipsum dolor amet swine biltong pork loin tail shoulder short loin. Filet mignon spare ribs chuck, kevin ribeye tail pancetta. Biltong salami landjaeger jowl. Ham turkey biltong, swine kielbasa alcatra doner shank rump picanha chuck.

Short loin frankfurter prosciutto tail, bresaola boudin flank picanha ham hock kevin. Beef turducken pork belly cupim. Capicola alcatra prosciutto strip steak brisket turkey chuck shank jerky picanha ground round. Turkey sirloin pork loin picanha t-bone ribeye.
Bacon ipsum dolor amet swine biltong pork loin tail shoulder short loin. Filet mignon spare ribs chuck, kevin ribeye tail pancetta. Biltong salami landjaeger jowl. Ham turkey biltong, swine kielbasa alcatra doner shank rump picanha chuck.

Short loin frankfurter prosciutto tail, bresaola boudin flank picanha ham hock kevin. Beef turducken pork belly cupim. Capicola alcatra prosciutto strip steak brisket turkey chuck shank jerky picanha ground round. Turkey sirloin pork loin picanha t-bone ribeye.
	`,
	verify: func(stdin []byte) (bool, error) {
		return true, nil
	},
}
var IntroTour = Content{
	Title: "Hello Mars",
	Text: `
	how this works
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
