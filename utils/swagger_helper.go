package utils

// giant bara description, and by bara I mean capybara

// Description for swagger documentation.
var SwaggerDescription string = "This is a documentation of the **acmcsuf API**.\n\n" +
	"---\n" +
	"# **acmcsuf-cli**\n" +
	"The command line interface tools provide a quick and easy way to test out the currently implemented routes.\n\n" +
	"## **csuf events [command]**\n" +
	"Use this command to interact with event-related features.\n\n" +
	"<h3>Available Commands</h3>\n" +
	"<ul>\n" +
	"  <li><strong>create</strong> – Create a new event</li>\n" +
	"  <li><strong>delete</strong> – Delete an event by ID</li>\n" +
	"  <li><strong>list</strong> – List all events</li>\n" +
	"  <li><strong>update</strong> – Update an event by ID</li>\n" +
	"</ul>\n\n" +
	"<h3>Flags</h3>\n" +
	"<ul>\n" +
	"  <li><code>-h</code>, <code>--help</code> – Show help for events</li>\n" +
	"</ul>\n\n" +
	"<h3>Available Commands & Flags</h3>\n" +
	"<ul>\n" +
	"  <li><strong>post</strong> – Post a new event\n" +
	"    <ul>\n" +
	"      <li><code>-d, --duration &lt;string&gt;</code> – Duration (03:04:05)</li>\n" +
	"      <li><code>-H, --host &lt;string&gt;</code> – Host</li>\n" +
	"      <li><code>-a, --isallday</code> – All-day event</li>\n" +
	"      <li><code>-l, --location &lt;string&gt;</code> – Location</li>\n" +
	"      <li><code>--port &lt;string&gt;</code> – Port (default 8080)</li>\n" +
	"      <li><code>-s, --startat &lt;string&gt;</code> – Start time</li>\n" +
	"      <li><code>--urlhost &lt;string&gt;</code> – URL host (default 127.0.0.1)</li>\n" +
	"      <li><code>-u, --uuid &lt;string&gt;</code> – UUID</li>\n" +
	"    </ul>\n" +
	"  </li>\n" +
	"  <li><strong>get</strong> – Get events\n" +
	"    <ul>\n" +
	"      <li><code>--host &lt;string&gt;</code> – Host (default 127.0.0.1)</li>\n" +
	"      <li><code>--id &lt;string&gt;</code> – Specific event ID</li>\n" +
	"      <li><code>--port &lt;string&gt;</code> – Port (default 8080)</li>\n" +
	"    </ul>\n" +
	"  </li>\n" +
	"  <li><strong>put</strong> – Update an event\n" +
	"    <ul>\n" +
	"      <li><code>--id &lt;string&gt;</code> – <strong>[REQUIRED]</strong> Event ID to update</li>\n" +
	"      <li><code>-d, --duration &lt;string&gt;</code> – End time</li>\n" +
	"      <li><code>-H, --host &lt;string&gt;</code> – Host</li>\n" +
	"      <li><code>-a, --isallday</code> – All-day event</li>\n" +
	"      <li><code>-l, --location &lt;string&gt;</code> – Location</li>\n" +
	"      <li><code>--port &lt;string&gt;</code> – Port (default 8080)</li>\n" +
	"      <li><code>-s, --startat &lt;string&gt;</code> – Start time</li>\n" +
	"      <li><code>--urlhost &lt;string&gt;</code> – URL host (default 127.0.0.1)</li>\n" +
	"      <li><code>-u, --uuid &lt;string&gt;</code> – UUID</li>\n" +
	"    </ul>\n" +
	"  </li>\n" +
	"  <li><strong>delete</strong> – Delete an event\n" +
	"    <ul>\n" +
	"      <li><code>--id &lt;string&gt;</code> – <strong>[REQUIRED]</strong> Event ID</li>\n" +
	"      <li><code>--host &lt;string&gt;</code> – Host (default 127.0.0.1)</li>\n" +
	"      <li><code>--port &lt;string&gt;</code> – Port (default 8080)</li>\n" +
	"    </ul>\n" +
	"  </li>\n" +
	"</ul>\n" +
	"<p><code>-h, --help</code> – Show help for events</p>\n\n" +
	"---\n" +
	"<h2><strong>csuf announcements [command]</strong></h2>\n" +
	"<p>Manage ACM CSUF's announcements through the CLI. Use <code>csuf announcements [command] --help</code> for more info on a specific command.</p>\n\n" +
	"<h3>Available Commands & Flags</h3>\n" +
	"<ul>\n" +
	"  <li><strong>post</strong> – Post a new announcement\n" +
	"    <ul>\n" +
	"      <li><code>-a, --announceat &lt;string&gt;</code> – Set this announcement's announce at</li>\n" +
	"      <li><code>-c, --channelid &lt;string&gt;</code> – Set this announcement's channel id</li>\n" +
	"      <li><code>--host &lt;string&gt;</code> – Custom host (default \"127.0.0.1\")</li>\n" +
	"      <li><code>-m, --messageid &lt;string&gt;</code> – Set this announcement's message id</li>\n" +
	"      <li><code>--port &lt;string&gt;</code> – Custom port (default \"8080\")</li>\n" +
	"      <li><code>--uuid &lt;string&gt;</code> – Set this announcement's id</li>\n" +
	"      <li><code>-v, --visibility &lt;string&gt;</code> – Set this announcement's visibility</li>\n" +
	"    </ul>\n" +
	"  </li>\n" +
	"  <li><strong>get</strong> – Get an announcement\n" +
	"    <ul>\n" +
	"      <li><code>--host &lt;string&gt;</code> – Custom host (default \"127.0.0.1\")</li>\n" +
	"      <li><code>--id &lt;string&gt;</code> – Get a specific announcement by its id</li>\n" +
	"      <li><code>--port &lt;string&gt;</code> – Custom port (default \"8080\")</li>\n" +
	"    </ul>\n" +
	"  </li>\n" +
	"  <li><strong>put</strong> – Update an existing announcement by its id\n" +
	"    <ul>\n" +
	"      <li><code>--id &lt;string&gt;</code> – <strong>[REQUIRED]</strong> Announcement id to update</li>\n" +
	"      <li><code>-a, --announceat &lt;string&gt;</code> – Change this announcement's announce at</li>\n" +
	"      <li><code>-c, --channelid &lt;string&gt;</code> – Change this announcement's discord channel id</li>\n" +
	"      <li><code>--host &lt;string&gt;</code> – Custom host (default \"127.0.0.1\")</li>\n" +
	"      <li><code>-m, --messageid &lt;string&gt;</code> – Change this announcement's discord message id</li>\n" +
	"      <li><code>--port &lt;string&gt;</code> – Custom port (default \"8080\")</li>\n" +
	"      <li><code>--uuid &lt;string&gt;</code> – Change this announcement's uuid</li>\n" +
	"      <li><code>-v, --visibility &lt;string&gt;</code> – Change this announcement's visibility</li>\n" +
	"    </ul>\n" +
	"  </li>\n" +
	"  <li><strong>delete</strong> – Delete an announcement\n" +
	"    <ul>\n" +
	"      <li><code>--id &lt;string&gt;</code> – <strong>[REQUIRED]</strong> Delete an announcement by its id</li>\n" +
	"      <li><code>--host &lt;string&gt;</code> – Custom host (default \"127.0.0.1\")</li>\n" +
	"      <li><code>--port &lt;string&gt;</code> – Custom port (default \"8080\")</li>\n" +
	"    </ul>\n" +
	"  </li>\n" +
	"</ul>\n" +
	"<p><code>-h, --help</code> – Show help for announcements</p>"
