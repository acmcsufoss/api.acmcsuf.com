This is a documentation of the **acmcsuf API**.

---
# **acmcsuf-cli**
The command line interface tools provide a quick and easy way to test out the currently implemented routes.

## **csuf events [command]**
Use this command to interact with event-related features.

<h3>Available Commands</h3>
<ul>
  <li><strong>create</strong> – Create a new event</li>
  <li><strong>delete</strong> – Delete an event by ID</li>
  <li><strong>list</strong> – List all events</li>
  <li><strong>update</strong> – Update an event by ID</li>
</ul>

<h3>Flags</h3>
<ul>
  <li><code>-h</code>, <code>--help</code> – Show help for events</li>
</ul>

<h3>Available Commands & Flags</h3>
<ul>
  <li><strong>post</strong> – Post a new event
    <ul>
      <li><code>-d, --duration &lt;string&gt;</code> – Duration (03:04:05)</li>
      <li><code>-H, --host &lt;string&gt;</code>     – Host</li>
      <li><code>-a, --isallday</code>                – All-day event</li>
      <li><code>-l, --location &lt;string&gt;</code> – Location</li>
      <li><code>--port &lt;string&gt;</code>         – Port (default 8080)</li>
      <li><code>-s, --startat &lt;string&gt;</code>  – Start time</li>
      <li><code>--urlhost &lt;string&gt;</code>      – URL host (default 127.0.0.1)</li>
      <li><code>-u, --uuid &lt;string&gt;</code>     – UUID</li>
    </ul>
  </li>
  <li><strong>get</strong> – Get events
    <ul>
      <li><code>--host &lt;string&gt;</code> – Host (default 127.0.0.1)</li>
      <li><code>--id &lt;string&gt;</code> – Specific event ID</li>
      <li><code>--port &lt;string&gt;</code> – Port (default 8080)</li>
    </ul>
  </li>
  <li><strong>put</strong> – Update an event
    <ul>
      <li><code>--id &lt;string&gt;</code> – <strong>[REQUIRED]</strong> Event ID to update</li>
      <li><code>-d, --duration &lt;string&gt;</code> – End time</li>
      <li><code>-H, --host &lt;string&gt;</code> – Host</li>
      <li><code>-a, --isallday</code> – All-day event</li>
      <li><code>-l, --location &lt;string&gt;</code> – Location</li>
      <li><code>--port &lt;string&gt;</code> – Port (default 8080)</li>
      <li><code>-s, --startat &lt;string&gt;</code> – Start time</li>
      <li><code>--urlhost &lt;string&gt;</code> – URL host (default 127.0.0.1)</li>
      <li><code>-u, --uuid &lt;string&gt;</code> – UUID</li>
    </ul>
  </li>
  <li><strong>delete</strong> – Delete an event
    <ul>
      <li><code>--id &lt;string&gt;</code> – <strong>[REQUIRED]</strong> Event ID</li>
      <li><code>--host &lt;string&gt;</code> – Host (default 127.0.0.1)</li>
      <li><code>--port &lt;string&gt;</code> – Port (default 8080)</li>
    </ul>
  </li>
</ul>
<p><code>-h, --help</code> – Show help for events</p>

---
<h2><strong>csuf announcements [command]</strong></h2>
<p>Manage ACM CSUF's announcements through the CLI. Use <code>csuf announcements [command] --help</code> for more info on a specific command.</p>

<h3>Available Commands & Flags</h3>
<ul>
  <li><strong>post</strong> – Post a new announcement
    <ul>
      <li><code>-a, --announceat &lt;string&gt;</code> – Set this announcement's announce at</li>
      <li><code>-c, --channelid &lt;string&gt;</code> – Set this announcement's channel id</li>
      <li><code>--host &lt;string&gt;</code> – Custom host (default "127.0.0.1")</li>
      <li><code>-m, --messageid &lt;string&gt;</code> – Set this announcement's message id</li>
      <li><code>--port &lt;string&gt;</code> – Custom port (default "8080")</li>
      <li><code>--uuid &lt;string&gt;</code> – Set this announcement's id</li>
      <li><code>-v, --visibility &lt;string&gt;</code> – Set this announcement's visibility</li>
    </ul>
  </li>
  <li><strong>get</strong> – Get an announcement
    <ul>
      <li><code>--host &lt;string&gt;</code> – Custom host (default "127.0.0.1")</li>
      <li><code>--id &lt;string&gt;</code> – Get a specific announcement by its id</li>
      <li><code>--port &lt;string&gt;</code> – Custom port (default "8080")</li>
    </ul>
  </li>
  <li><strong>put</strong> – Update an existing announcement by its id
    <ul>
      <li><code>--id &lt;string&gt;</code> – <strong>[REQUIRED]</strong> Announcement id to update</li>
      <li><code>-a, --announceat &lt;string&gt;</code> – Change this announcement's announce at</li>
      <li><code>-c, --channelid &lt;string&gt;</code> – Change this announcement's discord channel id</li>
      <li><code>--host &lt;string&gt;</code> – Custom host (default "127.0.0.1")</li>
      <li><code>-m, --messageid &lt;string&gt;</code> – Change this announcement's discord message id</li>
      <li><code>--port &lt;string&gt;</code> – Custom port (default "8080")</li>
      <li><code>--uuid &lt;string&gt;</code> – Change this announcement's uuid</li>
      <li><code>-v, --visibility &lt;string&gt;</code> – Change this announcement's visibility</li>
    </ul>
  </li>
  <li><strong>delete</strong> – Delete an announcement
    <ul>
      <li><code>--id &lt;string&gt;</code> – <strong>[REQUIRED]</strong> Delete an announcement by its id</li>
      <li><code>--host &lt;string&gt;</code> – Custom host (default "127.0.0.1")</li>
      <li><code>--port &lt;string&gt;</code> – Custom port (default "8080")</li>
    </ul>
  </li>
</ul>
<p><code>-h, --help</code> – Show help for announcements</p>
