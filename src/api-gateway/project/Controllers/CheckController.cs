using Microsoft.AspNetCore.Mvc;

namespace project.Controllers;

[ApiController]
[Route("[controller]")]
public class CheckController : ControllerBase
{
    [HttpGet(Name = "check")]
    public IActionResult Get()
    {
        return Ok("success");
    }
}