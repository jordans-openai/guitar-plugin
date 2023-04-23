## About

This plugin generates guitar tabs given a list of chords. It tries to select chords that are comfortable
to play, and that sound good together.

## Use It

You can curl the plugin directly to get ASCII art guitar tablature diagrams.

```shell
curl -s -X POST \
  -d '{"chords":[{"name":"Cmaj7"},{"name":"Dm"}]}' \
  https://guitar-plugin.fly.dev/chords \
  | jq -r '.tabs[]'
```

It is, however, designed to be installed as a plugin for ChatGPT. If plugins are enabled on your account, you
can install it by pointing ChatGPT to https://guitar-plugin.fly.dev/.

## Testing Locally

```shell
go run main.go
curl -s -X POST \
  -d '{"chords":[{"name":"Cmaj7"},{"name":"Dm"}]}' \
  http://localhost:8080/chords \
  | jq -r '.tabs[]'

```

## Acknowledgements

Many concepts and ideas were borrowed from [pcorey/chord](https://github.com/pcorey/chord/tree/master)!

## TODO
- [] the chords generated should follow logic to chain together well (voice leading, distance, etc)
- [] by default, the chords generated should be easy to play!
- 
## ideas
- each actual fingering should have an id (stateless) that encodes offset/fret/finger/name
- the api should accept excludes of ids, to let the user say "do it again, without that chord" or "give me an alternate"
- the api could return several tabs for each chords, each one with scores for reach, flow, invertedness, etc.
  they could be sorted with a default rule, and the the model may be smart enough to pick different ones based
  on user feedback. alternatively, each chord entry in the request could have optional fields to specify
  which attribute to sort on, and which direction to sort in.
