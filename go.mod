module github.com/pokt-network/poktroll

go 1.23.0

// replace (
// DEVELOPER_TIP: Uncomment to use a local copy of shannon-sdk for development purposes.
// github.com/pokt-network/shannon-sdk => ../shannon-sdk

// DEVELOPER_TIP: Uncomment to use a local copy of smt for development purposes.
// github.com/pokt-network/smt => ../smt
// github.com/pokt-network/smt/kvstore/badger => ../smt/kvstore/badger
// github.com/pokt-network/smt/kvstore/pebble => ../smt/kvstore/pebble
// )

replace nhooyr.io/websocket => github.com/coder/websocket v1.8.6

// replace broken goleveldb
replace github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7

require (
	cosmossdk.io/api v0.7.6
	cosmossdk.io/client/v2 v2.0.0-beta.8
	cosmossdk.io/core v0.11.1
	cosmossdk.io/depinject v1.1.0
	cosmossdk.io/errors v1.0.1
	cosmossdk.io/log v1.4.1
	cosmossdk.io/math v1.4.0
	cosmossdk.io/store v1.1.1
	cosmossdk.io/tools/confix v0.1.1
	cosmossdk.io/x/circuit v0.1.0
	cosmossdk.io/x/evidence v0.1.0
	cosmossdk.io/x/feegrant v0.1.0
	cosmossdk.io/x/upgrade v0.1.1
	github.com/athanorlabs/go-dleq v0.1.0
	github.com/bufbuild/buf v1.34.0
	github.com/cometbft/cometbft v0.38.12
	github.com/cosmos/cosmos-db v1.1.1
	github.com/cosmos/cosmos-proto v1.0.0-beta.5
	github.com/cosmos/cosmos-sdk v0.50.13
	github.com/cosmos/gogoproto v1.7.0
	github.com/cosmos/ibc-go/modules/capability v1.0.0
	github.com/cosmos/ibc-go/v8 v8.2.0
	github.com/go-kit/kit v0.13.0
	github.com/gogo/status v1.1.0
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/websocket v1.5.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	github.com/hashicorp/go-metrics v0.5.3
	github.com/pokt-network/ring-go v0.1.0
	// TODO_IMPROVE: Whenever we update a protobuf in the `poktroll` repo, we need to:
	// 1. Merge in the update PR (and it's generated outputs) into `poktroll` main.
	// 2. Update the `poktroll` sha in the `shannon-sdk` to reflect the new dependency.
	// 3. Update the `shannon-sdk` sha in the `poktroll` repo (here).
	// This is creating a circular dependency whereby exporting the protobufs into a separate
	// repo is the first obvious idea, but has to be carefully considered, automated, and is not
	// a hard blocker.
	github.com/pokt-network/shannon-sdk v0.0.0-20240907012836-7172ed278f8b
	github.com/pokt-network/smt v0.13.0
	github.com/pokt-network/smt/kvstore/pebble v0.0.0-20240822175047-21ea8639c188
	github.com/prometheus/client_golang v1.20.1
	github.com/regen-network/gocuke v1.1.0
	github.com/rs/zerolog v1.33.0
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.19.0
	github.com/stretchr/testify v1.10.0
	go.uber.org/multierr v1.11.0
	golang.org/x/crypto v0.31.0
	golang.org/x/exp v0.0.0-20240707233637-46b078467d37 // indirect
	golang.org/x/sync v0.10.0
	golang.org/x/text v0.21.0
	golang.org/x/tools v0.23.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240814211410-ddb44dafa142
	google.golang.org/grpc v1.67.1
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.4.0
	google.golang.org/protobuf v1.35.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	cosmossdk.io/x/tx v0.13.8
	github.com/jhump/protoreflect v1.16.0
	github.com/mitchellh/mapstructure v1.5.0
	go.uber.org/mock v0.5.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.34.2-20240508200655-46a4cf4ba109.2 // indirect
	buf.build/gen/go/bufbuild/registry/connectrpc/go v1.16.2-20240710174705-f2077dee5ad4.1 // indirect
	buf.build/gen/go/bufbuild/registry/protocolbuffers/go v1.34.2-20240710174705-f2077dee5ad4.2 // indirect
	cloud.google.com/go v0.112.1 // indirect
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	cloud.google.com/go/iam v1.1.6 // indirect
	cloud.google.com/go/storage v1.38.0 // indirect
	connectrpc.com/connect v1.16.2 // indirect
	connectrpc.com/otelconnect v0.7.0 // indirect
	cosmossdk.io/collections v0.4.0 // indirect
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.1 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/DataDog/datadog-go v3.2.0+incompatible // indirect
	github.com/DataDog/zstd v1.5.5 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/aws/aws-sdk-go v1.44.224 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/bgentry/speakeasy v0.1.1-0.20220910012023-760eaf8b6816 // indirect
	github.com/bits-and-blooms/bitset v1.8.0 // indirect
	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.4 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.2 // indirect
	github.com/bufbuild/protocompile v0.14.0 // indirect
	github.com/bufbuild/protoplugin v0.0.0-20240323223605-e2735f6c31ee // indirect
	github.com/bufbuild/protovalidate-go v0.6.3 // indirect
	github.com/bufbuild/protoyaml-go v0.1.9 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/cockroachdb/apd/v2 v2.0.2 // indirect
	github.com/cockroachdb/apd/v3 v3.2.1 // indirect
	github.com/cockroachdb/errors v1.11.3 // indirect
	github.com/cockroachdb/fifo v0.0.0-20240606204812-0bbfbd93a7ce // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/pebble v1.1.2 // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/cometbft/cometbft-db v0.11.0 // indirect
	github.com/containerd/continuity v0.4.2 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.15.1 // indirect
	github.com/cosmos/btcutil v1.0.5 // indirect
	github.com/cosmos/go-bip39 v1.0.0 // indirect
	github.com/cosmos/gogogateway v1.2.0 // indirect
	github.com/cosmos/iavl v1.2.2 // indirect
	github.com/cosmos/ics23/go v0.11.0 // indirect
	github.com/cosmos/ledger-cosmos-go v0.14.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/creachadair/atomicfile v0.3.1 // indirect
	github.com/creachadair/tomledit v0.0.24 // indirect
	github.com/cucumber/common/messages/go/v19 v19.1.2 // indirect
	github.com/cucumber/gherkin/go/v26 v26.2.0 // indirect
	github.com/cucumber/messages/go/v21 v21.0.1 // indirect
	github.com/cucumber/tag-expressions/go/v5 v5.0.6 // indirect
	github.com/danieljoos/wincred v1.2.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/dgraph-io/badger/v2 v2.2007.4 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/cli v27.0.3+incompatible // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/docker v27.1.1+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.8.2 // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/dvsekhvalnov/jose2go v1.6.0 // indirect
	github.com/emicklei/dot v1.6.2 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/felixge/fgprof v0.9.4 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-chi/chi/v5 v5.1.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/validator/v10 v10.15.5 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gofrs/uuid/v5 v5.2.0 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.2.4 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/google/cel-go v0.20.1 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-containerregistry v0.20.0 // indirect
	github.com/google/orderedcode v0.0.1 // indirect
	github.com/google/pprof v0.0.0-20240625030939-27f56978b8b0 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.3 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-getter v1.7.5 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-plugin v1.5.2 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/hdevalence/ed25519consensus v0.1.0 // indirect
	github.com/huandu/skiplist v1.2.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/improbable-eng/grpc-web v0.15.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jdx/go-netrc v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/pgzip v1.2.6 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/linxGnu/grocksdb v1.8.14 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/manifoldco/promptui v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/oasisprotocol/curve25519-voi v0.0.0-20230904125328-1f23a7beb09a // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.27.8 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/petermattis/goid v0.0.0-20231207134359-e60b3f734c67 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/profile v1.7.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d // indirect
	github.com/tendermint/go-amino v0.16.0 // indirect
	github.com/tidwall/btree v1.7.0 // indirect
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/vbatts/tar-split v0.11.5 // indirect
	github.com/zondax/hid v0.9.2 // indirect
	github.com/zondax/ledger-go v0.14.3 // indirect
	go.etcd.io/bbolt v1.3.10 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.49.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.53.0 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.22.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/mod v0.19.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/oauth2 v0.22.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/term v0.27.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/api v0.171.0 // indirect
	google.golang.org/genproto v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240930140551-af27646dc61f // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.5.1 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
	pgregory.net/rapid v1.1.0 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
