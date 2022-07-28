#!/usr/bin/env python3

from pathlib import Path
import os
import subprocess
import sys
import re
import argparse

def cleanup_files(base_path, test_only):
    cleaned_files = []
    for path in find_offending_dirs(base_path):
        package_name = extract_package_name(path)
        if package_name:
            new_path = rename_offending_directory(path, test_only)
            for f in replace_package_in_source_files( \
                    new_path, package_name, test_only):
                cleaned_files.append(os.path.join(new_path, f))
        else:
            print('Could not extract package name from' + path)

    return cleaned_files


# find any directories that have a '-' in them
def find_offending_dirs(base_path):
    offending_paths = []
    offending_pat = re.compile(r'.*-.*')

    for root, subdirs, files in os.walk(base_path):

        if offending_pat.match(root):
            offending_paths.append(root)

    return offending_paths

# extract a package name from a directory
def extract_package_name(offending_path):
    m = re.match(r'[a-zA-Z0-9/]*/(.*-[^/]*)', offending_path)
    if m:
        return m.group(1).replace("-", "_")
    return None

def replace_package_in_source_files(parent_path, package_name, test_only):
    replaced_files = []
    for file in os.listdir(parent_path):
        fqpn = os.path.join(parent_path, file)
        if file.endswith('.scala'):
            replaced_files.append(file)
            if not test_only:
                with open(fqpn, 'r') as rof:
                    contents = rof.read()

                new_contents = contents.replace( \
                        "__PACKAGE__", "package " + package_name)

                with open(fqpn, 'w') as wf:
                    wf.write(new_contents)

    return replaced_files

def rename_offending_directory(path, test_only):
    if not test_only:
        os.rename(path, path.replace("-", "_"))

    return path.replace("-", "_")

arger = argparse.ArgumentParser( \
        description='spiffy up files for leetup enjoyment')

arger.add_argument('--base', required=True, \
        help='Base path to start cleanup from (required)')
arger.add_argument('--open', action="store_true", \
        help='Open cleaned up files in a remote neovim session')
arger.add_argument('--test', action="store_true", \
        help='Do not do anything, just test and report')

ns = arger.parse_args()

if ns.test:
    print("---Test-mode---")

if not ns.base:
    arger.print_help()
    exit(1)

new_files = cleanup_files(ns.base, ns.test)

if not new_files:
    print("No files were found to clean")
    exit(0)

if ns.test:
    for f in new_files:
        print("Cleaned (Test) " + f)

if ns.open:
    nvr_cmd = ['nvr', '--remote', '-p' ]
    for nf in new_files:
        nvr_cmd.append(nf)

    if ns.test:
        print("would have run: " + nvr_cmd)
        exit(0)

    print("Running: " + ' '.join(nvr_cmd))
    subprocess.run(nvr_cmd)
