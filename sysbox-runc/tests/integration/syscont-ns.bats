#!/usr/bin/env bats

load helpers

function setup() {
	teardown_busybox
	setup_busybox
}

function teardown() {
	teardown_busybox
}

@test "syscont: uses all namespaces" {

	runc run -d --console-socket "$CONSOLE_SOCKET" test_busybox
	[ "$status" -eq 0 ]

	# For each ns, check that the sys container's init process is in a
	# different namespace than the test script.

	for nsType in cgroup ipc mnt net pid user uts; do
		syscont_ns=$(runc exec test_busybox ls -l /proc/1/ns | grep -i "$nsType" | cut -d":" -f3)
		[ "$status" -eq 0 ]
		test_ns=$(ls -l /proc/self/ns | grep -i "$nsType" | cut -d":" -f3)
		[ "$status" -eq 0 ]
		[ "$syscont_ns" != "$test_ns" ]
	done
}

@test "syscont: unshare" {

	runc run -d --console-socket "$CONSOLE_SOCKET" test_busybox
	[ "$status" -eq 0 ]

	# check that unshare(2) works inside a system container
	runc exec test_busybox sh -c "unshare -i -m -n -p -u -f --mount-proc echo 1 > /dev/null"
	[ "$status" -eq 0 ]

	# TODO: test that nsenter(2) also works
}
