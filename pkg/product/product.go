package product

type Category string

var (
	PhysicalMediaCategory            Category = "physical_media"
	BookCategory                     Category = "book"
	MemberAssociationCategory        Category = "member_association"
	MemberAssociationUpgradeCategory Category = "member_association_upgrade"
	VideoCategory                    Category = "video"
)

type Product struct {
	Name     string
	Category Category
	IsMember bool
}
