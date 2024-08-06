variable "team" {
  description = "Team name"
  type        = string
}
variable "members" {
  description = "Team members"
  type        = set(string)
}
variable "maintainers" {
  description = "Team maintainers"
  type        = set(string)
}