#!/bin/bash

# Call the binary ./sshpf and store the output in a variable
ssh_command=$(~/bin/sshpfbin)

# Execute the stored command as a bash command
bash -c "$ssh_command"
