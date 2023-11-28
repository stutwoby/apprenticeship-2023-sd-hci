using DependencyInjectionDemo.Services;
using LiteDB;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddRazorPages();

builder.Services.AddSingleton<IDatabaseService, LiteDatabaseService>(); 
builder.Services.AddScoped<IMenuService, MenuService>();     // I live as long as we're in the same "scoped" context
builder.Services.AddTransient<IIntermediaryService, IntemediaryService>();  // Any new request, new service

var app = builder.Build();

// Configure the HTTP request pipeline.
if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Error");
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
    app.UseHsts();
}

app.UseHttpsRedirection();
app.UseStaticFiles();

app.UseRouting();

app.UseAuthorization();

app.MapRazorPages();

app.Run();
