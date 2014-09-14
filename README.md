# codns

This code is borrowed from github.com/piger/nasello and adapted to easily work with
consul.io.

A very simple DNS proxy server capable of routing client queries to
different remote servers based on pattern matching.

The code is mostly based on [Go-DNS][go-dns] examples.

[go-dns]: http://miek.nl/projects/godns/

Warning: this is alpha software and should be used with caution.

## Getting Started

Running `sudo ./codns` you'll have a proxy dns running on port 127.0.0.1:53 TCP and UDP
that resolves consul database at 127.0.0.1:8600 and the rest of the dns space.

### Getting codns

The latest release is available at [Github][github-src]

[github-src]: https://github.com/jmcarbo/codns

### Installing from source

You can install nasello from source:

	go get github.com/jmcarbo/codns
	go install github.com/jmcarbo/codns/codns

### Configuration format

The configuration file is a JSON document which must contains a
`filters` list with one or more `pattern` dictionaries; each `filter`
must contain a *FQDN* DNS name as the `pattern` and a list of one or
more remote DNS servers to forward the query to. For *reverse lookups*
the `in-addr.arpa` domain must be used in the pattern definition.

The "." `pattern` specifies a default remote resolver.

### Example

The following configuration example specifies three `filters`:

- `*.consul` will be resolved by Codns (localhost:8600)
- `*.example.com` will be resolved by OpenDNS (208.67.222.222, etc.)
- `10.0.24.*` will also be resolved by OpenDNS
- all the other queries will be forwarded to Google DNS (8.8.8.8,
  etc.)

`codns.json`:

	{
		"filters": [
				{
						"pattern": "consul.",
						"addresses": [ "127.0.0.1:8600" ]
				},
				{
						"pattern": "example.com.",
						"addresses": [ "208.67.222.222", "208.67.220.220" ]
				},
				{
						"pattern": "24.0.10.in-addr.arpa.",
						"addresses": [ "208.67.222.222", "208.67.220.220" ]
				},
				{
						"pattern": ".",
						"addresses": [ "8.8.8.8", "8.8.4.4" ]
				}
		]
	}

## License

codns is under the MIT license. See the [LICENSE][license] file for
details.

[license]: https://github.com/jmcarbo/codns/blob/master/LICENSE
