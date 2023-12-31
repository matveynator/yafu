#!/bin/bash
app_name="yafu"
version="0.1-002"
git_root_path=`git rev-parse --show-toplevel`

go mod init ${app_name}
go mod download
go mod vendor
go mod tidy

execution_file="${app_name}"

echo "Performing tests on all modules..."
go test ./...
if [ $? != "0" ] 
	then
		echo "Tests on all modules failed."
		echo "Press any key to continue compilation or CTRL+C to abort."
		read
	else 
		echo "Tests on all modules passed."
fi

cd ${git_root_path}/scripts;

mkdir -p ${git_root_path}/binaries/${version};

rm -f ${git_root_path}/binaries/latest; 

cd ${git_root_path}/binaries; ln -s ${version} latest; cd ${git_root_path}/scripts;

for target_os in linux freebsd netbsd openbsd aix android illumos ios solaris plan9 darwin dragonfly windows;
#for target_os in darwin;
do
	for arch in "amd64" "386" "arm" "arm64" "mips64" "mips64le" "mips" "mipsle" "ppc64" "ppc64le" "riscv64" "s390x" "wasm"
	do
		target_os_name=${target_os}
		[ "$target_os" == "windows" ] && execution_file="${app_name}.exe"
		[ "$target_os" == "darwin" ] && target_os_name="mac"
		
		mkdir -p ../binaries/${version}/${target_os_name}/${arch}

		GOOS=${target_os} GOARCH=${arch} go build -ldflags "-X ${app_name}/pkg/config.CompileVersion=${version}" -o ../binaries/${version}/${target_os_name}/${arch}/${execution_file} ../${app_name}.go 2> /dev/null


		if [ "$?" != "0" ]
		#if compilation failed - remove folders - else copy config file.
		then
		  rm -rf ../binaries/${version}/${target_os_name}/${arch}
		else
		  echo "GOOS=${target_os} GOARCH=${arch} go build -ldflags "-X ${app_name}/pkg/config.CompileVersion=${version}" -o ../binaries/${version}/${target_os_name}/${arch}/${execution_file} ../${app_name}.go"
      chmod +x ../binaries/${version}/${target_os_name}/${arch}/${execution_file}
      cp ../candidates.txt ../binaries/${version}/${target_os_name}/${arch}/
		fi
	done
done

#optional: publish to internet:
rsync -avP ../binaries/* files@files.matveynator.ru:/home/files/public_html/${app_name}/
