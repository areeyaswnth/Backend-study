using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
namespace WebApi.Controllers
{
    [ApiController]
    [Route("[controller]")]
    
    public class DemoController:ControllerBase
    {
        [Route("Countries")]
        [HttpGet]
        public List<string> GetCourntries(string match)
        {
            var items = new List<string>
            {
                "Thailand",
                "USA",
                "Korea",
                "Japan"
            };
            if(match == null){
                return items;
            }

            var results = items.Where(item => item.Contains(match));

            return results.ToList();
        }
    }
}
