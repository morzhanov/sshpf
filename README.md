# sshpf
Performs port forwarding of apis via ssh using a config

## Usage

1. Create config and store it in /bin
2. Store sshpf.sh file in /bin
3. Build and store app in the /bin/sshpfbin
4. Modify PATH and use tool

```bash
sshpf <?env>
```

## Config

Config file should be stored in the `~/bin/sshpf_config.yaml`

## Bash file

Bash file should be stored in the `~/bin/sshpf.sh`

## Binary file

Binary file should be stored in the `~/bin/sshpfbin`

## Building and adding to PATH (fish)

1. **Compile your Go script into a binary**:

   Compile your script into a binary by running:

    ```bash
    go build .
    ```

2. **Move the binary to a directory in your PATH**:

   Create a directory for your custom binaries. You can create a `bin` directory in your home directory:

    ```bash
    mkdir ~/bin
    ```

   Move the `sshpf` binary to this directory:

    ```bash
    mv sshpf ~/bin
    ```

3. **Add the directory to your PATH in Fish shell**:

   Open your Fish shell configuration file `~/.config/fish/config.fish`:

    ```bash
    nano ~/.config/fish/config.fish
    ```

   Add the following line to this file to include the `bin` directory in your PATH:

    ```fish
    set -gx PATH $HOME/bin $PATH
   
   alias sshpf "bash ~/bin/sshpf.sh"
    ```

   Save the changes to the configuration file.

4. **Reload Fish configuration**:

   After saving the changes, reload your Fish configuration to apply the changes:

    ```bash
    source ~/.config/fish/config.fish
    ```

5. **Test your setup**:

   You should now be able to run `sshpf` from anywhere in your terminal, and it should execute your `sshpf` script with the specified task name.
