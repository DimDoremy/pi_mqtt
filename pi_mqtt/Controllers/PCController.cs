using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using pi_mqtt.Dal;
using pi_mqtt.Entity;

namespace pi_mqtt.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class PCController : ControllerBase
    {
        private readonly MonitorContext _dbContext;

        public PCController()
        {
            _dbContext = new MonitorContext();
        }

        [HttpGet("GetPcResult")]
        public JsonResult GetPcResult(int id)
        {
            try
            {
                var ans = _dbContext.PC.Where(p => p.Id == id).ToList();
                return new JsonResult(new Response(ans, 0));
            }
            catch (Exception e)
            {
                return new JsonResult(new Response(e.Message, 1));
            }
        }

        [HttpDelete("DeletePcResult")]
        public JsonResult DeletePcResult(int id)
        {
            try
            {
                var del = _dbContext.PC.Where(p => p.Id == id);
                _dbContext.RemoveRange(del);
                _dbContext.SaveChanges();
                return new JsonResult(new Response("OK", 0));
            }
            catch (Exception e)
            {
                return new JsonResult(new Response(e.Message, 1));
            }
        }
    }
}
