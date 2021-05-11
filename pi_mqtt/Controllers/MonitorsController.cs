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
    public class MonitorsController : ControllerBase
    {
        private readonly MonitorContext _dbContext;

        public MonitorsController()
        {
            _dbContext = new MonitorContext();
        }

        [HttpGet("GetMonitorResult")]
        public JsonResult GetMonitorResult(int id)
        {
            try
            {
                var ans = _dbContext.Monitor.Where(p => p.Id == id).ToList();
                return new JsonResult(new Response(ans, 0));
            }
            catch (Exception e)
            {
                return new JsonResult(new Response(e.Message, 1));
            }
        }

        [HttpDelete("DeleteMonitorResult")]
        public JsonResult DeleteMonitorResult(int id)
        {
            try
            {
                var del = _dbContext.Monitor.Where(p => p.Id == id);
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
