locals {
  github_owner = "thong0905Org"

  yaml_file = yamldecode(file("${path.module}/input.yaml"))
  users = toset(concat([
    for team in local.yaml_file :
      concat(
        try(team.member, []),
        try(team.maintainer, []))
    ]...))
  team_memberships = {
    for team in local.yaml_file :
      team.name => {
        members     = toset(try(team.member, []))
        maintainers = toset(try(team.maintainer, []))
    }}
}