package product

type Category string

var (
	PhysicalMediaCategory            Category = "physical_media"
	BookCategory                     Category = "book"
	memberAssociationCategory        Category = "member_association"
	memberAssociationUpgradeCategory Category = "member_association_upgrade"
	VideoCategory                    Category = "video"
)

type Product struct {
	Name     string
	Category Category
}
