function translateToEnglish() {
    // Resumo Profissional
    document.getElementById('resume-text').textContent = 'Developer addicted to Programming with solid experience in frontend and backend development. Currently, I am solidifying my knowledge in Go and backend development, with a strong emphasis on efficient and scalable programming practices.';
    
    // Habilidades Técnicas
    document.getElementById('skills').innerHTML = `
        <li><strong>Programming Languages:</strong> JavaScript, HTML, CSS, GO, TypeScript</li>
        <li><strong>Frameworks and Libraries:</strong> React, Vue, Axios, GorillaMux(GO), Nextjs</li>
        <li><strong>Development Tools:</strong> Git, Webpack, PostgreSQL</li>
        <li><strong>Other Skills:</strong> REST APIs, AJAX</li>
        <li><strong>Currently Learning:</strong> HTTP Request Caching and Axios Interceptors</li>
    `;
    
    // Experiência Profissional
    document.getElementById('job1-title').textContent = 'User Registration Project';
    document.getElementById('job1-company').textContent = 'Junior Developers Group';
    document.getElementById('job1-divider').textContent = '____________________________________';
    document.getElementById('job1-tasks').innerHTML = `
        <li>Developed reusable components in React for use in various web applications.</li>
        <li>Collaborated with designers to implement responsive user interfaces.</li>
        <li>Implemented API calls to integrate dynamic functionalities.</li>
        <li>Fixed bugs and optimized web page performance, and participated in the basic construction of the backend.</li>
    `;
    document.getElementById('job2-title').textContent = 'Social Network';
    document.getElementById('job2-company').textContent = 'Sole Developer';
    document.getElementById('job2-divider').textContent = '____________________________________';
    document.getElementById('job2-tasks').innerHTML = `
        <li><strong>Social Network Development:</strong> Created a social network using React for the frontend and Go for the backend. Implemented features such as JWT authentication, post creation, and retrieval.</li>
        <li><strong>API Integration:</strong> Developed RESTful endpoints for handling posts and managing users, with secure and efficient access control.</li>
        <li><strong>Design and Navigation:</strong> Created and styled a responsive and intuitive user interface, including profile pages and navigation system.</li>
        <li><strong>Optimization and Performance:</strong> Implemented caching techniques to improve HTTP request performance and ensure a fast user experience.</li>
        <li><strong>Database:</strong> Utilized PostgreSQL for data management, focusing on security and data integrity.</li>
        <li><a href="https://github.com/brenod/my-social-network">Social Network Repository</a></li>
    `;
    
    // Educação
    document.getElementById('education-info').innerHTML = '<strong>Technical Administration Course Integrated with High School - IFES</strong><br>Federal Institute of Espírito Santo in Venda Nova do Imigrante<br>[2020 - 2022]';
    document.getElementById('certifications').innerHTML = '<strong>Courses and Certifications</strong>';
    document.getElementById('certifications-list').innerHTML = '<li>Cod3r Technology School: Programming Logic, JavaScript, React, Vue, among others.</li>';
    
    // Idiomas
    document.getElementById('languages-list').innerHTML = `
        <li><strong>Portuguese:</strong> Native</li>
        <li><strong>English:</strong> Advanced</li>
    `;
}
