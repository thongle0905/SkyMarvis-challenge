resource "github_team" "team" {
  name        = var.team
  description = "Team ${var.team}"
}

resource "github_team_membership" "member" {
  for_each = var.members
  team_id  = github_team.team.id
  username = each.value
  role     = "member"
}

resource "github_team_membership" "maintainer" {
  for_each = var.maintainers
  team_id  = github_team.team.id
  username = each.value
  role     = "maintainer"
}