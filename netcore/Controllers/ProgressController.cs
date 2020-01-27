using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace netcore.Controllers
{
    [Route("player/[controller]")]
    [ApiController]
    public class ProgressController : ControllerBase
    {
        public static Dictionary<int, int> database = new Dictionary<int, int>();

        // GET api/values/5
        [HttpGet("{id}")]
        public ActionResult<dynamic> Get(int id)
        {
            if (!database.ContainsKey(id)){
                database.Add(id, 0);
            }
            return new { progress = database[id] };
        }

        // PUT api/values/5
        [HttpPatch("{id}/{position}")]
        public void Patch(int id, int position)
        {
            if (!database.ContainsKey(id)){
                database.Add(id, 0);
            }
            database[id] = position;
        }
    }
}
