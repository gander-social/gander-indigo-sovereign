package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	comatproto "github.com/gander-social/gander-indigo-sovereign/api/atproto"
	appgndr "github.com/gander-social/gander-indigo-sovereign/api/gndr"
	lexutil "github.com/gander-social/gander-indigo-sovereign/lex/util"
	"github.com/gander-social/gander-indigo-sovereign/util"
	"github.com/gander-social/gander-indigo-sovereign/util/cliutil"

	cli "github.com/urfave/cli/v2"
)

var gndrCmd = &cli.Command{
	Name:  "gndr",
	Usage: "sub-commands for gndr-specific endpoints",
	Subcommands: []*cli.Command{
		gndrFollowCmd,
		gndrListFollowsCmd,
		gndrPostCmd,
		gndrGetFeedCmd,
		gndrLikeCmd,
		gndrDeletePostCmd,
		gndrActorGetSuggestionsCmd,
		gndrNotificationsCmd,
	},
}

var gndrFollowCmd = &cli.Command{
	Name:      "follow",
	Usage:     "create a follow relationship (auth required)",
	Flags:     []cli.Flag{},
	ArgsUsage: `<user>`,
	Action: func(cctx *cli.Context) error {
		xrpcc, err := cliutil.GetXrpcClient(cctx, true)
		if err != nil {
			return err
		}

		user := cctx.Args().First()

		follow := appgndr.GraphFollow{
			LexiconTypeID: "gndr.app.graph.follow",
			CreatedAt:     time.Now().Format(time.RFC3339),
			Subject:       user,
		}

		resp, err := comatproto.RepoCreateRecord(context.TODO(), xrpcc, &comatproto.RepoCreateRecord_Input{
			Collection: "gndr.app.graph.follow",
			Repo:       xrpcc.Auth.Did,
			Record:     &lexutil.LexiconTypeDecoder{Val: &follow},
		})
		if err != nil {
			return err
		}

		fmt.Println(resp.Uri)

		return nil
	},
}

var gndrListFollowsCmd = &cli.Command{
	Name:      "list-follows",
	Usage:     "print list of follows for account",
	ArgsUsage: `[actor]`,
	Action: func(cctx *cli.Context) error {
		xrpcc, err := cliutil.GetXrpcClient(cctx, false)
		if err != nil {
			return err
		}

		user := cctx.Args().First()
		if user == "" {
			user = xrpcc.Auth.Did
		}

		ctx := context.TODO()
		resp, err := appgndr.GraphGetFollows(ctx, xrpcc, user, "", 100)
		if err != nil {
			return err
		}

		for _, f := range resp.Follows {
			fmt.Println(f.Did, f.Handle)
		}

		return nil
	},
}

var gndrPostCmd = &cli.Command{
	Name:      "post",
	Usage:     "create a post record",
	ArgsUsage: `<text>`,
	Action: func(cctx *cli.Context) error {
		xrpcc, err := cliutil.GetXrpcClient(cctx, true)
		if err != nil {
			return err
		}

		auth := xrpcc.Auth

		text := strings.Join(cctx.Args().Slice(), " ")

		resp, err := comatproto.RepoCreateRecord(context.TODO(), xrpcc, &comatproto.RepoCreateRecord_Input{
			Collection: "gndr.app.feed.post",
			Repo:       auth.Did,
			Record: &lexutil.LexiconTypeDecoder{Val: &appgndr.FeedPost{
				Text:      text,
				CreatedAt: time.Now().Format(util.ISO8601),
			}},
		})
		if err != nil {
			return fmt.Errorf("failed to create post: %w", err)
		}

		fmt.Println(resp.Cid)
		fmt.Println(resp.Uri)

		return nil
	},
}

func prettyPrintPost(p *appgndr.FeedDefs_FeedViewPost, uris bool) {
	fmt.Println(strings.Repeat("-", 60))
	rec := p.Post.Record.Val.(*appgndr.FeedPost)
	fmt.Printf("%s (%s)", p.Post.Author.Handle, rec.CreatedAt)
	if uris {
		fmt.Println(" -- ", p.Post.Uri)
	} else {
		fmt.Println(":")
	}
	fmt.Println(rec.Text)
}

var gndrGetFeedCmd = &cli.Command{
	Name:  "get-feed",
	Usage: "fetch gndr feed",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "count",
			Value: 100,
		},
		&cli.StringFlag{
			Name:  "author",
			Usage: "specify handle of user to list their authored feed",
		},
		&cli.BoolFlag{
			Name:  "raw",
			Usage: "print out feed in raw json",
		},
		&cli.BoolFlag{
			Name:  "uris",
			Usage: "include URIs in pretty print output",
		},
	},
	Action: func(cctx *cli.Context) error {
		xrpcc, err := cliutil.GetXrpcClient(cctx, true)
		if err != nil {
			return err
		}

		ctx := context.TODO()

		raw := cctx.Bool("raw")

		uris := cctx.Bool("uris")

		author := cctx.String("author")
		if author != "" {
			if author == "self" {
				author = xrpcc.Auth.Did
			}

			tl, err := appgndr.FeedGetAuthorFeed(ctx, xrpcc, author, "", "", false, 99)
			if err != nil {
				return err
			}

			for i := len(tl.Feed) - 1; i >= 0; i-- {
				it := tl.Feed[i]
				if raw {
					jsonPrint(it)
				} else {
					prettyPrintPost(it, uris)
				}
			}
		} else {
			algo := "reverse-chronological"
			tl, err := appgndr.FeedGetTimeline(ctx, xrpcc, algo, "", int64(cctx.Int("count")))
			if err != nil {
				return err
			}

			for i := len(tl.Feed) - 1; i >= 0; i-- {
				it := tl.Feed[i]
				if raw {
					jsonPrint(it)
				} else {
					prettyPrintPost(it, uris)
				}
			}
		}

		return nil

	},
}

var gndrActorGetSuggestionsCmd = &cli.Command{
	Name:      "actor-get-suggestions",
	ArgsUsage: "[author]",
	Action: func(cctx *cli.Context) error {
		xrpcc, err := cliutil.GetXrpcClient(cctx, true)
		if err != nil {
			return err
		}

		ctx := context.TODO()

		author := cctx.Args().First()
		if author == "" {
			author = xrpcc.Auth.Did
		}

		resp, err := appgndr.ActorGetSuggestions(ctx, xrpcc, "", 100)
		if err != nil {
			return err
		}

		b, err := json.MarshalIndent(resp.Actors, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(b))

		return nil

	},
}

var gndrLikeCmd = &cli.Command{
	Name:      "like",
	Usage:     "create gndr 'like' record",
	ArgsUsage: "<post>",
	Action: func(cctx *cli.Context) error {
		xrpcc, err := cliutil.GetXrpcClient(cctx, true)
		if err != nil {
			return err
		}

		arg := cctx.Args().First()

		parts := strings.Split(arg, "/")
		if len(parts) < 3 {
			return fmt.Errorf("invalid post uri: %q", arg)
		}
		rkey := parts[len(parts)-1]
		collection := parts[len(parts)-2]
		did := parts[2]

		fmt.Println(did, collection, rkey)
		ctx := context.TODO()
		resp, err := comatproto.RepoGetRecord(ctx, xrpcc, "", collection, did, rkey)
		if err != nil {
			return fmt.Errorf("getting record: %w", err)
		}

		out, err := comatproto.RepoCreateRecord(ctx, xrpcc, &comatproto.RepoCreateRecord_Input{
			Collection: "gndr.app.feed.like",
			Repo:       xrpcc.Auth.Did,
			Record: &lexutil.LexiconTypeDecoder{
				Val: &appgndr.FeedLike{
					CreatedAt: time.Now().Format(util.ISO8601),
					Subject:   &comatproto.RepoStrongRef{Uri: resp.Uri, Cid: *resp.Cid},
				},
			},
		})
		if err != nil {
			return fmt.Errorf("creating like failed: %w", err)
		}
		_ = out
		return nil

	},
}

var gndrDeletePostCmd = &cli.Command{
	Name:      "delete-post",
	ArgsUsage: `<rkey>`,
	Action: func(cctx *cli.Context) error {
		xrpcc, err := cliutil.GetXrpcClient(cctx, true)
		if err != nil {
			return err
		}

		rkey := cctx.Args().First()

		if rkey == "" {
			return fmt.Errorf("must specify rkey of post to delete")
		}

		schema := "gndr.app.feed.post"
		if strings.Contains(rkey, "/") {
			parts := strings.Split(rkey, "/")
			schema = parts[0]
			rkey = parts[1]
		}

		_, err = comatproto.RepoDeleteRecord(context.TODO(), xrpcc, &comatproto.RepoDeleteRecord_Input{
			Repo:       xrpcc.Auth.Did,
			Collection: schema,
			Rkey:       rkey,
		})
		return err
	},
}

var gndrNotificationsCmd = &cli.Command{
	Name:  "notifs",
	Usage: "fetch gndr notifications (requires auth)",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		ctx := context.TODO()

		xrpcc, err := cliutil.GetXrpcClient(cctx, true)
		if err != nil {
			return err
		}

		notifs, err := appgndr.NotificationListNotifications(ctx, xrpcc, "", 50, false, nil, "")
		if err != nil {
			return err
		}

		for _, n := range notifs.Notifications {
			b, err := json.Marshal(n)
			if err != nil {
				return err
			}

			fmt.Println(string(b))
		}

		return nil
	},
}
