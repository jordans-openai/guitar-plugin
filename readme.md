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
- [ ] bug: muted and open strings are not handled separately. this ruins the voicing scoring and creates misleading tabs
- [ ] bug: the finger locations on many chords just look wrong...
- [ ] the chords generated should follow logic to chain together well (voice leading, distance, etc)
- [ ] the chords API should accept hints for each chord, like `{"name":"Cmaj7", "hints": {"inversion": 2}}`
- [ ] render images instead of ASCII
