#version 420
#extension GL_ARB_explicit_uniform_location : enable
#extension GL_ARB_enhanced_layouts : enable

layout(location = 0) in vec3 vp;

layout(binding = 1) uniform PerspectiveBlock {
    mat4 project;
    mat4 camera;
    mat4 model;
};

void main() {
    gl_Position = project * camera * model * vec4(vp, 1);
}