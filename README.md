# Neurahive Client
Go implementation for client to interact with storage nodes in Neurahive network.

# SDK

Application could use a `node/Client` instance to interact with storage node via JSON RPC. Especially, use `Client.KV()` for **KV** operations.

# CLI
Run `go build` under the root folder to compile the executable binary.

**Global options**
```
      --gas-limit uint     Custom gas limit to send transaction
      --gas-price uint     Custom gas price to send transaction
  -h, --help               help for neurahive-client
      --log-force-color    Force to output colorful logs
      --log-level string   Log level (default "info")
```

**Deploy contract**

```
./neurahive-client deploy --url <blockchain_rpc_endpoint> --key <private_key> --bytecode <bytecode_hex_or_json_file>
```

**Generate test file**

To generate a file for test purpose, especially with a fixed file size or random file size (without `--size` option):

```
./neurahive-client gen --size <file_size_in_bytes>
```

**Upload file**
```
./neurahive-client upload --url <blockchain_rpc_endpoint> --contract <neurahive_contract_address> --key <private_key> --node <storage_node_rpc_endpoint> --file <file_path>
```

**Download file**
```
./neurahive-client download --node <storage_node_rpc_endpoint> --root <file_root_hash> --file <output_file_path>
```

To download file from multiple storage nodes **in parallel**, `--node` option supports to specify multiple comma separated URLs, e.g. `url1,url2,url3`.

If you want to verify the **merkle proof** of downloaded segment, please specify `--proof` option.