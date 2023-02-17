using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Kurdi.ECommerce.Inventory.Core.Entities.CategoryAggregate;

[Table(name: "categories")]
public class Category
{
    [Key]
    [Column(name: "name")]
    public string Name { get; set; }
    [Column(name: "is_parent")]
    public bool IsParent { get; set; }
    [Column(name: "parent")]
    public string ParentName { get; set; }
    [ForeignKey(name: "ParentName")]
    public Category Parent { get; set; }
    public List<CategoryDetails> CategoryDetails { get; set; } = new List<CategoryDetails>();
    [Column(name: "activation")]
    public bool Activation { get; set; }
    public TimeStamps TimeStamps { get; set; }
}