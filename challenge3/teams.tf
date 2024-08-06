# Add users to org
resource "github_membership" "org_member" {
  for_each = local.users
  username = each.value
  role     = "member"
}

module "team_memberships" {
  for_each    = local.team_memberships
  source      = "./modules/team_memberships"
  team        = each.key
  members     = each.value.members
  maintainers = each.value.maintainers
}