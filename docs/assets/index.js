(function() {
    let docute = new Docute({
        target: '#docute',
        detectSystemDarkTheme: true,
        darkThemeToggler: true,
        sidebar: Sidebar,
        nav: Navbar,
        footer: `
          <div style="border-top:1px solid var(--border-color);padding-top:30px;margin: 40px 0;color:#999999;font-size: .9rem;">
          &copy; ${new Date().getFullYear()} Developed with ♥ by <a href="http://fluidtech.in" target="_blank">Fluid Tech</a>. Released under MIT license.
          </div>
          `,
    })
})();