package all

import (
	// base_gcmessages protocol
	_ "github.com/paralin/go-dota2/protocol/base_gcmessages"
	// econ_gcmessages protocol
	_ "github.com/paralin/go-dota2/protocol/econ_gcmessages"
	// econ_shared_enums protocol
	_ "github.com/paralin/go-dota2/protocol/econ_shared_enums"
	// dota_client_enums protocol
	_ "github.com/paralin/go-dota2/protocol/dota_client_enums"
	// dota_gcmessages_client protocol
	_ "github.com/paralin/go-dota2/protocol/dota_gcmessages_client"
	// dota_gcmessages_client_chat protocol
	_ "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_chat"
	// dota_gcmessages_client_guild protocol
	_ "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_guild"
	// dota_gcmessages_client_match_management protocol
	_ "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_match_management"
	// dota_gcmessages_client_team protocol
	_ "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_team"
	// dota_gcmessages_client_fantasy protocol
	_ "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_fantasy"
	// dota_gcmessages_client_watch protocol
	_ "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_watch"
	// dota_gcmessages_client_tournament protocol
	_ "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_tournament"
	// dota_gcmessages_common protocol
	_ "github.com/paralin/go-dota2/protocol/dota_gcmessages_common"
	// dota_gcmessages_common_match_management protocol
	_ "github.com/paralin/go-dota2/protocol/dota_gcmessages_common_match_management"
	// dota_gcmessages_msgid protocol
	_ "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
	// dota_shared_enums protocol
	_ "github.com/paralin/go-dota2/protocol/dota_shared_enums"
	// gcsdk_gcmessages protocol
	_ "github.com/paralin/go-dota2/protocol/gcsdk_gcmessages"
	// gcsystemmsgs protocol
	_ "github.com/paralin/go-dota2/protocol/gcsystemmsgs"
	// steammessages protocol
	_ "github.com/paralin/go-dota2/protocol/steammessages"
)

// Packages is a list of all packages.
var Packages []string = []string{
	"github.com/paralin/go-dota2/protocol/base_gcmessages",
	"github.com/paralin/go-dota2/protocol/econ_gcmessages",
	"github.com/paralin/go-dota2/protocol/econ_shared_enums",
	"github.com/paralin/go-dota2/protocol/dota_client_enums",
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client",
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_chat",
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_guild",
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_match_management",
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_team",
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_fantasy",
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_watch",
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_tournament",
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_common",
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_common_match_management",
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid",
	"github.com/paralin/go-dota2/protocol/dota_shared_enums",
	"github.com/paralin/go-dota2/protocol/gcsdk_gcmessages",
	"github.com/paralin/go-dota2/protocol/gcsystemmsgs",
	"github.com/paralin/go-dota2/protocol/steammessages",
}
