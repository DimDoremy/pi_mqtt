using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.InteropServices;
using System.Threading.Tasks;
using pi_mqtt.Entity;

namespace pi_mqtt.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class RpioController : ControllerBase
    {
        [DllImport("libRpioRpc.dll", EntryPoint = "High")]
        private static extern void High(byte[] address, uint pin);

        [DllImport("libRpioRpc.dll", EntryPoint = "Low")]
        private static extern void Low(byte[] address, uint pin);

        [DllImport("libRpioRpc.dll", EntryPoint = "ReadPin")]
        private static extern int ReadPin(byte[] address, uint pin);

        [DllImport("libRpioRpc.dll", EntryPoint = "PullUp")]
        private static extern void PullUp(byte[] address, uint pin);

        [DllImport("libRpioRpc.dll", EntryPoint = "PullDown")]
        private static extern void PullDown(byte[] address, uint pin);


        [HttpPost("SetHigh")]
        public JsonResult SetHigh([FromBody] Rpio rpio)
        {
            try
            {
                var ad = System.Text.Encoding.UTF8.GetBytes(rpio.Address + ":" + rpio.Port.ToString());
                High(ad, rpio.Pin);
                return new JsonResult(new Response("OK", 0));
            }
            catch (Exception e)
            {
                return new JsonResult(new Response(e.Message, 1));
            }
        }

        [HttpPost("SetLow")]
        public JsonResult SetLow([FromBody] Rpio rpio)
        {
            try
            {
                var ad = System.Text.Encoding.UTF8.GetBytes(rpio.Address + ":" + rpio.Port.ToString());
                Low(ad, rpio.Pin);
                return new JsonResult(new Response("OK", 0));
            }
            catch (Exception e)
            {
                return new JsonResult(new Response(e.Message, 1));
            }
        }

        [HttpGet("ReadPin")]
        public JsonResult ReadPin(string address, uint pin)
        {
            try
            {
                var ad = System.Text.Encoding.UTF8.GetBytes(address);
                var readPin = ReadPin(ad, pin);
                return new JsonResult(new Response(readPin == 1 ? "High" : "Low", 0));
            }
            catch (Exception e)
            {
                return new JsonResult(new Response(e.Message, 1));
            }
        }

        [HttpPost("PullUp")]
        public JsonResult PullUp([FromBody] Rpio rpio)
        {
            try
            {
                var ad = System.Text.Encoding.UTF8.GetBytes(rpio.Address + ":" + rpio.Port.ToString());
                PullUp(ad, rpio.Pin);
                return new JsonResult(new Response("OK", 0));
            }
            catch (Exception e)
            {
                return new JsonResult(new Response(e.Message, 1));
            }
        }


        [HttpPost("PullDown")]
        public JsonResult PullDown([FromBody] Rpio rpio)
        {
            try
            {
                var ad = System.Text.Encoding.UTF8.GetBytes(rpio.Address + ":" + rpio.Port.ToString());
                PullDown(ad, rpio.Pin);
                return new JsonResult(new Response("OK", 0));
            }
            catch (Exception e)
            {
                return new JsonResult(new Response(e.Message, 1));
            }
        }
    }
}
