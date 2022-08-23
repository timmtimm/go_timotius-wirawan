# Version Control and Branch Management (Git)

## Resume Materi

### Git

Git merupakan software berbasis Version Control System (VCS) yang bertugas untuk mencatat perubahan seluruh file atau repository suatu project. Developer software biasa menggunakan Git untuk distributed revision (VCS terdistribusi), hal ini bertujuan untuk menyimpan database tidak hanya ke satu tempat. Namun semua orang yang terlibat dalam penyusunan kode dapat menyimpan database ini.

### Command

Berikut ini perintah-perintah Git yang akan sering digunakan

| Command                                           | Description                                                        |
|---------------------------------------------------|--------------------------------------------------------------------|
| [config](https://git-scm.com/docs/git-config)     | Get and set repository or global options                           |
| [init](https://git-scm.com/docs/git-init)         | Create an empty Git repository or reinitialize an existing one     |
| [remote](https://git-scm.com/docs/git-remote)     | Manage set of tracked repositories                                 |
| [push](https://git-scm.com/docs/git-push)         | Update remote refs along with associated objects                   |
| [clone](https://git-scm.com/docs/git-clone)       | Clone a repository into a new directory                            |
| [add](https://git-scm.com/docs/git-add)           | Add file contents to the index                                     |
| [commit](https://git-scm.com/docs/git-commit)     | Record changes to the repository                                   |
| [status](https://git-scm.com/docs/git-status)     | Show the working tree status                                       |
| [diff](https://git-scm.com/docs/git-diff)         | Show changes between commits, commit and working tree, etc         |
| [stash](https://git-scm.com/docs/git-stash)       | Stash the changes in a dirty working directory away                |
| [log](https://git-scm.com/docs/git-log)           | Show commit logs                                                   |
| [checkout](https://git-scm.com/docs/git-checkout) | Switch branches or restore working tree files                      |
| [fetch](https://git-scm.com/docs/git-fetch)       | Download objects and refs from another repository                  |
| [pull](https://git-scm.com/docs/git-pull)         | Fetch from and integrate with another repository or a local branch |
| [branch](https://git-scm.com/docs/git-branch)     | List, create, or delete branches                                   |
| [merge](https://git-scm.com/docs/git-merge)       | Join two or more development histories together                    |

### Workflow Collaboration

Untuk proses development dimana akan berkolaborasi dalam tim, kita tidak bisa hanya bekerja dalam satu branch sehingga perlu dibuat beberapa branch agar kolaborasi dapat berjalan dengan optimal.

Langkah-langkah yang perlu dilakukan yaitu:
1. Buat branch master dari branch development
2. Hindari direct edit ke branch development
3. Merge branch feature hanya ke branch development
4. Merge branch development ke branch master jika development selesai

Adapun beberapa hal yang perlu diperhatikan yaitu:
- Jika terjadi conflict, resolve dengan satu computer, gunakan plugin untuk membandingkan perubahan
- Jika terlalu banyak branch, diperbolehkan untuk menghapus branch yang sudah lama dan tidak terpakai